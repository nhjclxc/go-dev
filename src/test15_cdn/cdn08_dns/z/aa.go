package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	host       = "https://www.51dns.com"
	domainID   = "92274924"
	apiKey     = "c7722149110b7492a2e5cf1d8f3f966b"
	apiSecret  = "ecb4ff0e877a83292b9f35067e9ae673"
	timeoutSec = 5
)

type APIResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RecordCount int `json:"recordCount"`
		Data        []struct {
			RecordID string `json:"recordID"`
			Value    string `json:"value"`
		} `json:"data"`
	} `json:"data"`
}

// ================= 签名 =================

func sign(params map[string]string, secret string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	for i, k := range keys {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(params[k])
	}

	sum := md5.Sum([]byte(b.String() + secret))
	return hex.EncodeToString(sum[:])
}

// ================= HTTP POST =================

func post(path string, params map[string]string) ([]byte, error) {
	form := url.Values{}
	for k, v := range params {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", host+path, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		Timeout: time.Second * timeoutSec,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

// ================= 查询 =================

func searchRecord(hostname string) (*APIResp, error) {
	params := map[string]string{
		"apiKey":    apiKey,
		"domainID":  domainID,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		"host":      hostname,
		"regex":     "2",
		"page":      "1",
		"pageSize":  "1",
	}
	params["hash"] = sign(params, apiSecret)

	fmt.Println(params)
	return nil, nil

	//body, err := post("/api/record/search/", params)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var resp APIResp
	//err = json.Unmarshal(body, &resp)
	//return &resp, err
}

// ================= 更新 =================

func updateRecord(recordID, ip string) error {
	params := map[string]string{
		"apiKey":    apiKey,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		"domainID":  domainID,
		"recordID":  recordID,
		"newvalue":  ip,
	}
	params["hash"] = sign(params, apiSecret)

	fmt.Println(params)

	return nil
}

func main() {
	updateRecord("", "")
}

/*

map[apiKey:c7722149110b7492a2e5cf1d8f3f966b domainID:92274924 hash:be2aa29a4503abfa4ef90dd15ce315dd newvalue: recordID: timestamp:1769140912]


array:4 [\n  \"domainID\" => \"jscloud.com\"\n  \"timestamp\" => \"1769138725\"\n  \"apiKey\" => \"bbada2fd5a0dbced2b0b13366356147b\"\n  \"hash\" => \"7885223aa794156fed317ade5852bedc\"\n]\

bbada2fd5a0dbced2b0b13366356147b 4a54615ca4d3e38ad960bc41e60066c9

apiKey=bbada2fd5a0dbced2b0b13366356147b&domainID=jscloud.com×tamp=17691387254a54615ca4d3e38ad960bc41e60066c9



apiKey=c7722149110b7492a2e5cf1d8f3f966b&domain=51dns.com&timestamp=1521005892&secret=ecb4ff0e877a83292b9f35067e9ae673
1b6fe97150a3ff11a71702966a486d5e

apiKey=c7722149110b7492a2e5cf1d8f3f966b&domain=51dns.com&timestamp=1521005892ecb4ff0e877a83292b9f35067e9ae673
1196a23f87995b219f6aeb83dfb716a6


*/
