package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"sort"
	"time"
)

func GetHash(param map[string]string, secert string) string {
	var keyList []string
	for k, _ := range param {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	var hashString string
	for _, key := range keyList {
		if hashString == "" {
			hashString += key + "=" + param[key]
		} else {
			hashString += "&" + key + "=" + param[key]
		}
	}

	m := md5.New()
	m.Write([]byte(hashString + secert))
	cipherStr := m.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

type PostDomainBody struct {
	ApiKey    string `json:"apiKey"`
	Timestamp string `json:"timestamp"`
	Hash      string `json:"hash"`
	DomainID  string `json:"domainID"`
}

func getDomainID(paramMap map[string]string) {
	client := resty.New()
	paramMap["domainID"] = "jscloud.com"
	hash := GetHash(paramMap, secert)
	paramMap["hash"] = hash

	fmt.Println(paramMap)

	resp, err := client.R().
		SetQueryParams(paramMap).
		//SetBody(body).
		Post("https://www.51dns.com/api/domain/getsingle/")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.StatusCode())

	var diRes DomainInfoResponse
	err = json.Unmarshal(resp.Body(), &diRes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(diRes)
}

type DomainInfoResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    DomainInfo `json:"data"`
}
type DomainInfo struct {
	DomainID  int    `json:"domainID"`  // 域名ID
	Domain    string `json:"domain"`    // 域名
	State     int    `json:"state"`     // 状态：0-正常，2-暂停，3-管理员暂停
	Lock      int    `json:"lock"`      // 用户锁定：0-正常，1-用户锁定
	AdminLock int    `json:"adminLock"` // 管理员锁定：0-正常，1-管理员锁定
	ServiceID string `json:"serviceID"` // 服务套餐ID
	NSGroupID int    `json:"nsGroupID"` // NS组ID
}

func GetParsedRecord(paramMap map[string]string, domainID string) {
	client := resty.New()
	paramMap["domainID"] = domainID
	hash := GetHash(paramMap, secert)
	paramMap["hash"] = hash

	fmt.Println(paramMap)
	resp, err := client.R().
		SetQueryParams(paramMap).
		Get("https://www.51dns.com/api/record/getsingle/")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.StatusCode())
	fmt.Println(resp.String())
}

// bbada2fd5a0dbced2b0b13366356147b 4a54615ca4d3e38ad960bc41e60066c9
var apiKey string = "bbada2fd5a0dbced2b0b13366356147b"
var secert string = "4a54615ca4d3e38ad960bc41e60066c9"

func main() {
	ts := fmt.Sprintf("%d", time.Now().Unix())
	var param = map[string]string{}
	param["apiKey"] = apiKey
	param["domain"] = "51dns.com"
	param["timestamp"] = ts

	//getDomainID(param)

	fmt.Println()

	domainID := "142740793"
	var param2 = map[string]string{}
	param2["apiKey"] = apiKey
	param2["domain"] = "51dns.com"
	param2["timestamp"] = ts
	GetParsedRecord(param2, domainID)

}

// 获取单个域名 https://www.51dns.com/document/api/74/31.html
// 获取单个解析记录 https://www.51dns.com/document/api/4/49.html
// apiKey=c7722149110b7492a2e5cf1d8f3f966b&domain=51dns.com&timestamp=1521005892ecb4ff0e877a83292b9f35067e9ae673
