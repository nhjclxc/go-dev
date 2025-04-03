package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Num       int
}

func main() {

	/*
		<Person>
			<FirstName>Laura</FirstName>
			<LastName>Lynn</LastName>
		</Person>

		如同 json 包一样，也有 xml.Marshal() 和 xml.Unmarshal() 从 XML 中编码和解码数据；
		但这个更通用，可以从文件中读取和写入（或者任何实现了 io.Reader 和 io.Writer 接口的类型）
		和 JSON 的方式一样，XML 数据可以序列化为结构，或者从结构反序列化为 XML 数据；

		encoding/xml 包实现了一个简单的 XML 解析器（SAX），用来解析 XML 数据内容。下面的例子说明如何使用解析器：
	*/

	//text := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"

	p := Person{
		FirstName: "Luo",
		LastName:  "Chao",
		Num:       18,
	}

	// 序列化
	xmlByte, err1 := xml.Marshal(p)
	xmlStr := string(xmlByte)
	fmt.Println(err1)
	fmt.Println(xmlByte)
	fmt.Println(xmlStr)

	// 反序列化
	p2 := Person{}
	fmt.Println(p2)
	err2 := xml.Unmarshal(xmlByte, &p2)
	fmt.Println(err2)
	fmt.Println(p2)
	fmt.Println(p2.Num)
	fmt.Println(p2.FirstName)
	fmt.Println(p2.LastName)

	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	decoder := xml.NewDecoder(inputReader)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
	}

}
