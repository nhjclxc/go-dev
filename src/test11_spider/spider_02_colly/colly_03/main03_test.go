package colly_03

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"testing"
)

// 保存【https://rollcall.com/factbase/transcript/】上面每一个演讲内容为json
// https://m.w3cschool.cn/colly/colly-examples-factbase.html

func TestMain03(t *testing.T) {

	listC := colly.NewCollector()

	//detailC := listC.Clone()
	detailC := colly.NewCollector(
		colly.Async(true),
	)

	var detailMap map[string][]Detail = make(map[string][]Detail)

	listC.OnResponse(func(resp *colly.Response) {
		var listResultData ListResultData
		json.Unmarshal(resp.Body, &listResultData)

		//fmt.Println("dataItems = ", listResultData)

		for _, data := range listResultData.Data {
			ctx := colly.NewContext()
			ctx.Put("recordTitle", data.RecordTitle)
			detailC.Request("GET", data.FactbaseURL, nil, ctx, nil)
			//break
		}
	})

	detailC.OnRequest(func(request *colly.Request) {
		fmt.Println("开始请求 recordTitle", request.Ctx.Get("recordTitle"))
	})

	detailC.OnHTML("div.flex.gap-4.py-2 div.w-full", func(e *colly.HTMLElement) {
		recordTitle := e.Request.Ctx.Get("recordTitle")
		//fmt.Println("recordTitle", recordTitle)

		name := e.ChildText("h2.text-md")
		content := e.ChildText("div.flex-auto.text-md.text-gray-600.leading-loose")
		//fmt.Println("找到了 div", name, content)

		detailMap[recordTitle] = append(detailMap[recordTitle],
			Detail{
				Name:    name,
				Content: content,
			})

	})

	detailC.OnResponse(func(response *colly.Response) {
		fmt.Println("结束请求 recordTitle", response.Ctx.Get("recordTitle"))
	})

	listC.Visit("https://rollcall.com/wp-json/factbase/v1/search?media=&type=&sort=date&location=all&place=all&page=1&format=json&person=trump")
	detailC.Wait()

	// 保存
	fmt.Println("len(detailMap) = ", len(detailMap))

	for key, val := range detailMap {
		save2Json("./json", key, val)
	}

}

func save2Json(path, key string, val []Detail) {

	// 方法1: Marshal + WriteFile
	data, err := json.MarshalIndent(val, "", "  ") // 美化缩进
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}

	err = os.WriteFile(path+"/"+key+".json", data, 0644)
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}

	fmt.Printf("保存成功 %s/%s.json \n", path, key)

	//file, err := os.Create(key + ".json")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//encoder := json.NewEncoder(file)
	//encoder.SetIndent("", "  ") // 美化格式
	//if err := encoder.Encode(val); err != nil {
	//	panic(err)
	//}
}

type Detail struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ListResultData struct {
	Meta MetaStruct   `json:"meta"`
	Data []DataStruct `json:"data"`
}
type MetaStruct struct {
	RecordsMatched int `json:"records_matched"`
	RecordsTotal   int `json:"records_total"`
	Page           int `json:"page"`
	TotalPages     int `json:"total_pages"`
	ResultsPerPage int `json:"results_per_page"`
	Milliseconds   int `json:"milliseconds"`
}

type DataStruct struct {
	Candidate     string `json:"candidate"`
	Date          string `json:"date"`
	DocumentID    string `json:"document_id"`
	FactbaseURL   string `json:"factbase_url"`
	ImageURL      string `json:"image_url"`
	RecordType    string `json:"record_type"`
	MediaType     string `json:"media_type"`
	Place         string `json:"place"`
	RecordTitle   string `json:"record_title"`
	Slug          string `json:"slug"`
	Source        string `json:"source"`
	SpeakerID     string `json:"speaker_id"`
	Type          string `json:"type"`
	URL           string `json:"url"`
	Version       string `json:"version"`
	VersionNumber int    `json:"version_number"`
	VideoURL      string `json:"video_url"`
	Location      struct {
		City      string `json:"city"`
		State     string `json:"state"`
		StateCode string `json:"state_code"`
		Country   string `json:"country"`
	} `json:"location"`
	Topics []struct {
		Topic         string      `json:"topic"`
		Score         float64     `json:"score"`
		Magnitude     int         `json:"magnitude"`
		Code          int         `json:"code"`
		URL           string      `json:"url"`
		WikidataID    interface{} `json:"wikidata_id"`
		WikidataLabel interface{} `json:"wikidata_label"`
		Source        string      `json:"source"`
	} `json:"topics"`
	Moderation []struct {
		Category       string  `json:"category"`
		ModerationFlag bool    `json:"moderation_flag"`
		Score          float64 `json:"score"`
		Source         string  `json:"source"`
	} `json:"moderation"`
	Stresslens interface{} `json:"stresslens"`
	Document   struct {
		Entities []struct {
			Entity        string      `json:"entity"`
			Type          string      `json:"type"`
			Score         float64     `json:"score"`
			WikidataID    interface{} `json:"wikidata_id"`
			WikidataLabel interface{} `json:"wikidata_label"`
		} `json:"entities"`
		Speakers []struct {
			Speaker                   string  `json:"speaker"`
			SpeakerID                 string  `json:"speaker_id"`
			TotalTime                 string  `json:"total_time"`
			TotalSeconds              int     `json:"total_seconds"`
			SpeakerSeconds            int     `json:"speaker_seconds"`
			SpeakerSecondsPct         float64 `json:"speaker_seconds_pct"`
			TotalSentences            int     `json:"total_sentences"`
			SpeakerSentences          int     `json:"speaker_sentences"`
			SpeakerSentencesPct       float64 `json:"speaker_sentences_pct"`
			TotalWords                int     `json:"total_words"`
			SpeakerWords              int     `json:"speaker_words"`
			SpeakerWordsPct           float64 `json:"speaker_words_pct"`
			SentimentLoughranMcdonald float64 `json:"sentiment_loughran_mcdonald"`
			SentimentHarvardIv        float64 `json:"sentiment_harvard_iv"`
			SentimentVader            float64 `json:"sentiment_vader"`
		} `json:"speakers"`
		Sentiment []struct {
			Label string  `json:"label"`
			Score float64 `json:"score"`
			Text  string  `json:"text"`
		} `json:"sentiment"`
		Topics []struct {
			Topic         string      `json:"topic"`
			Score         float64     `json:"score"`
			Magnitude     interface{} `json:"magnitude"`
			Code          int         `json:"code"`
			URL           string      `json:"url"`
			WikidataID    string      `json:"wikidata_id"`
			WikidataLabel string      `json:"wikidata_label"`
			Source        string      `json:"source"`
		} `json:"topics"`
		Moderation []struct {
			Category       string  `json:"category"`
			ModerationFlag bool    `json:"moderation_flag"`
			Score          float64 `json:"score"`
			Source         string  `json:"source"`
		} `json:"moderation"`
		Stresslens []interface{} `json:"stresslens"`
		WordCount  interface{}   `json:"word_count"`
		Duration   string        `json:"duration"`
	} `json:"document"`
	VideoID int `json:"video_id"`
	Video   struct {
		VimeoStart string `json:"vimeo_start"`
	} `json:"video"`
}
