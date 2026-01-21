package tests

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

var (
	host    = "http://127.0.0.1:8080" // 注意: 固定使用该请求域
	uri     = "/api/v1/get"           // 注意：请求uri, 需要用真实uri替换
	ak      = "ak"                    // 注意：真实accessKey
	sk      = "sk"                    // 注意：真实secretKey
	version = "1.0"                   // 注意：版本号，当前支持固定传1.0
	service = "CDN"                   // 注意：服务号，当前支持固定传CDN
)

func genHMAC256(ciphertext, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(ciphertext))
	hmac := mac.Sum(nil)
	return hmac
}

func genAuthHeader(ak, sk string, ts int64, version, service string) map[string]string {
	sign := fmt.Sprintf("%s%d%s%s", service, ts, ak, sk)
	hmac := genHMAC256([]byte(sign), []byte(sk))
	authorization := hex.EncodeToString(hmac)

	header := make(map[string]string)
	header["X-WY-Authorization"] = fmt.Sprintf("HMAC-SHA256 Certificate=AccessKey/Timestamp/Service, Sign=%s", authorization)
	header["X-WY-AccessKey"] = ak
	header["X-WY-Timestamp"] = strconv.FormatInt(ts, 10)
	header["X-WY-version"] = version
	header["X-WY-service"] = service
	header["Content-Type"] = "application/json"
	return header
}

func TestName(t *testing.T) {

	// 需要根据请求方法是POST或者GET设置request,
	url := fmt.Sprintf("%s%s", host, uri)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	ts := time.Now().Unix()
	header := genAuthHeader(ak, sk, ts, version, service)
	for k, v := range header {
		req.Header.Set(k, v)
		fmt.Println(k, "-", v)
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("response body: %+v\n", string(body))
}

func TestName2(t *testing.T) {
	z, x, _ := GenerateCredential()

	fmt.Println(z)
	fmt.Println(x)

}

func GenerateAccessKey() string {
	u := uuid.New().String() // 36 位
	u = strings.ReplaceAll(u, "-", "")
	return "GKCDNAK_" + u[:16]
}

func GenerateAccessSecret() (string, error) {
	const length = 32 // 32 字节 = 256 bit
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GenerateCredential() (string, string, error) {
	sk, err := GenerateAccessSecret()
	if err != nil {
		return "", "", err
	}

	return GenerateAccessKey(), sk, nil
}
