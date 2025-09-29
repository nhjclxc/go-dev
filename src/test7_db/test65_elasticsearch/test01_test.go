package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"log"
	"net/http"

	"testing"
)

func Test0101(t *testing.T) {
	// 连接方式1：跳过证书认证，只适合开发环境

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200", // Docker 映射端口
		},
		Username: "elastic",
		Password: "NGN2iFwDf1-9*5wfLPd5",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 测试连接
	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res)

}

func Test0102(t *testing.T) {
	// 连接方式2：导入证书认证
	// 在'/Users/lxc20250729'路径下执行：“docker cp es01:/usr/share/elasticsearch/config/certs/http_ca.crt ." 将es的证书导出

	// 读取 CA 证书
	caCert, err := ioutil.ReadFile("/Users/lxc20250729/http_ca.crt")
	if err != nil {
		log.Fatalf("Error reading CA cert: %s", err)
	}

	// 加入证书池
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cfg := elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Username:  "elastic",
		Password:  "NGN2iFwDf1-9*5wfLPd5",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 测试连接
	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res)

}

func Test0103(t *testing.T) {
	getEsClient()
}

func getEsClient() *elasticsearch.Client {
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200", // Docker 映射端口
		},
		Username: "elastic",
		Password: "NGN2iFwDf1-9*5wfLPd5",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 测试连接
	res, err := esClient.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("elasticsearch connect successful !!!")

	return esClient
}
