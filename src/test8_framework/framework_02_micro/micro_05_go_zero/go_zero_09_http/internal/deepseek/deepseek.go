package deepseek

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {


	content := "You are a helpful assistant"

	//var msgChan chan []string = make(chan []string, 500)
	var msgChan chan string = make(chan string)
	var exitChan chan bool = make(chan bool)
	defer close(msgChan)
	defer close(exitChan)


	// 写数据
	go SendDeepSeek(exitChan, msgChan, content)

	// 读数据
	for {
		select {
		case <-exitChan:
			return
		case msg := <-msgChan:
			fmt.Println("读取:", msg)
		}
	}





}

func SendDeepSeek(exitChan chan bool, msgChan chan string, content string) {
	url := "https://api.deepseek.com/chat/completions"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(
		`{
			  "messages": [
				{
				  "content": "%s",
				  "role": "system"
				},
				{
				  "content": "Hi",
				  "role": "user"
				}
			  ],
			  "model": "deepseek-chat",
			  "frequency_penalty": 0,
			  "max_tokens": 2048,
			  "presence_penalty": 0,
			  "response_format": {
				"type": "text"
			  },
			  "stop": null,
			  "stream": true,
			  "stream_options": null,
			  "temperature": 1,
			  "top_p": 1,
			  "tools": null,
			  "tool_choice": "none",
			  "logprobs": false,
			  "top_logprobs": null
			}`, content))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	token := readToken("D:\\code\\go\\deepseekToken.txt")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	dataChip := DeepSeekDataChip{}
	for scanner.Scan() {
		// 第一个数据片
		// data: {"id":"e3420702-e3c5-4a1c-b39f-4de7f126bd24","object":"chat.completion.chunk","created":1749188623,"model":"deepseek-chat","_fingerprint":"fp_8802369eaa_prod0425fp8","choices":[{"index":0,"delta":{"role":"assistant","content":""},"logprobs":null,"finish_reason":null}]}
		// 中间的数据片
		// data: {"id":"e3420702-e3c5-4a1c-b39f-4de7f126bd24","object":"chat.completion.chunk","created":1749188623,"model":"deepseek-chat","system_fingerprint":"fp_8802369eaa_prod0425fp8","choices":[{"index":0,"delta":{"content":" 😊"},"logprobs":null,"finish_reason":null}]}
		// 最后一个数据片
		// data: {"id":"e3420702-e3c5-4a1c-b39f-4de7f126bd24","object":"chat.completion.chunk","created":1749188623,"model":"deepseek-chat","system_finge rprint":"fp_8802369eaa_prod0425fp8","choices":[{"index":0,"delta":{"content":""},"logprobs":null,"finish_reason":"stop"}],"usage":{"prompt_tokens":9,"completion_tokens":11,"total_tokens":20,"prompt_tokens_details":{"cached_tokens":0},"prompt_cache_hit_tokens":0,"prompt_cache_miss_tokens":9}}
		// 结束标记
		//data: [DONE]

		line := scanner.Text()
		//line = strings.Replace(line, "data: ", "", 1)
		if strings.HasPrefix(line, "data: ") {
			line = strings.TrimPrefix(line, "data: ")
		}
		if "[DONE]" == line {
			exitChan <- true
		}
		if "" == line {
			continue
		}

		err := json.Unmarshal([]byte(line), &dataChip)
		if err != nil {
			exitChan <- true
			return
		}
		content := dataChip.Choices[0].Delta.Content
		if "" == content {
			continue
		}

		//data := []string{fmt.Sprintf(content)}
		//fmt.Println("接收到数据:", data)
		msgChan <- fmt.Sprintf(content)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取数据出错:", err)
	}
}

type DeepSeekDataChip struct {
	Id                string `json:"id"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index int `json:"index"`
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason interface{} `json:"finish_reason"`
	} `json:"choices"`
}

func readToken(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return ""
	}

	// 输出内容（string 转换）
	//fmt.Println(string(data))

	return string(data)
}