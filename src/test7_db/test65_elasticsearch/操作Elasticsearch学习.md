
https://www.liwenzhou.com/posts/Go/go-elasticsearch/


【【【文件来源：https://github.com/guyan0319/golang_development_notes/blob/master/zh/4.6.md】】】

## go操作elasticsearch示例
这里我使用elasticsearch官方给的go语言包（[go-elasticsearch](https://github.com/elastic/go-elasticsearch)）
go-elasticsearch向前兼容，这意味着客户端支持与更大或同等次要版本的 Elasticsearch 通信。Elasticsearch 语言客户端仅向后兼容默认发行版，不提供任何保证。

- 包：https://github.com/elastic/go-elasticsearch

- Elasticsearch: 权威指南：https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html

### 环境介绍：
    版本
    Elasticsearch:v7.15


### 安装
go.mod 文件中添加
```
require github.com/elastic/go-elasticsearch/v8 main
```
或者
````
git clone --branch main https://github.com/elastic/go-elasticsearch.git $GOPATH/src/github.com/elastic/go-elasticsearch
````
### 示例：
新建 es.go 存入 es目录
````
package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net/http"
)

var EsClient *elasticsearch.Client

func init() {

	cfg := elasticsearch.Config{
		 Addresses: []string{
            "http://localhost:9200",
            },
	}
	var err error
	EsClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln("Failed to connect to es")
	}
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// idx 为空，默认随机唯一字符串
func Index(index, idx string, doc map[string]interface{}) {
	//index:="my_index_name_v1"
	res, err := EsClient.Info()
	fmt.Println(res, err)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	var buf bytes.Buffer
	//doc := map[string]interface{}{
	//	"title":   "中国",
	//	"content": "中国早日统一台湾",
	//	"time":    time.Now().Unix(),
	//	"date":    time.Now(),
	//}
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		fmt.Println(err, "Error encoding doc")
		return
	}
	res, err = EsClient.Index(
		index,                              // Index name
		&buf,                               // Document body
		EsClient.Index.WithDocumentID(idx), // Document ID
		// Document ID
		EsClient.Index.WithRefresh("true"), // Refresh
	)
	//res, err = EsClient.Create(index, idx, &buf)
	if err != nil {
		fmt.Println(err, "Error create response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
	log.Println(res)
}

//struct 类型允许使用更实际的方法，您可以在其中创建一个新结构，将请求配置作为字段，并使用上下文和客户端作为参数调用 Do() 方法：
func IndexEspi(index, idx string, doc map[string]interface{}) {
	//index:="my_index_name_v1"
	res, err := EsClient.Info()
	fmt.Println(res, err)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	var buf bytes.Buffer
	//doc := map[string]interface{}{
	//	"title":   "中国",
	//	"content": "中国早日统一台湾",
	//	"time":    time.Now().Unix(),
	//	"date":    time.Now(),
	//}
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		fmt.Println(err, "Error encoding doc")
		return
	}

	req := esapi.IndexRequest{
		Index:      index,  // Index name
		Body:       &buf,   // Document body
		DocumentID: idx,    // Document ID
		Refresh:    "true", // Refresh
	}

	res, err = req.Do(context.Background(), EsClient)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	fmt.Println(res.String())
	log.Println(res)
}
func Search(index string, query map[string]interface{}) {
	res, err := EsClient.Info()
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	//fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": title,
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = EsClient.Search(
		EsClient.Search.WithContext(context.Background()),
		EsClient.Search.WithIndex(index),
		EsClient.Search.WithBody(&buf),
		EsClient.Search.WithTrackTotalHits(true),
		EsClient.Search.WithFrom(0),
		EsClient.Search.WithSize(10),
		EsClient.Search.WithSort("time:desc"),
		EsClient.Search.WithPretty(),
	)
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

//删除 index 根据 索引名 id

func Delete(index, idx string) {
	//index:="my_index_name_v1"
	res, err := EsClient.Info()
	fmt.Println(res, err)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	res, err = EsClient.Delete(
		index, // Index name
		idx,   // Document ID
		EsClient.Delete.WithRefresh("true"),
	)
	if err != nil {
		fmt.Println(err, "Error create response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
	log.Println(res)
}
func DeleteByQuery(index []string, query map[string]interface{}) {
	res, err := EsClient.Info()
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	//fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": title,
	//		},
	//	},
	//	},
	//}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = EsClient.DeleteByQuery(
		index,
		&buf,
	)
	if err != nil {
		fmt.Println(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())

}
func SearchEsapiSql(query map[string]interface{}) {
	jsonBody, _ := json.Marshal(query)
	req := esapi.SQLQueryRequest{Body: bytes.NewReader(jsonBody)}
	res, _ := req.Do(context.Background(), EsClient)
	defer res.Body.Close()
	fmt.Println(res.String())
}
func SearchHttp(method, url string, query map[string]interface{}) {
	jsonBody, _ := json.Marshal(query)
	req, _ := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	req.Header.Add("Content-type", "application/json")
	res, err := EsClient.Perform(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	fmt.Println(buf.String())
}

````
新建 main.go

````
package main
import "demo/es"
func main() {
	index := "my_index_name_v4"
	//创建索引并设置映射
	//query := map[string]interface{}{
	//	"mappings": map[string]interface{}{
	//		"properties": map[string]interface{}{
	//			"title": map[string]interface{}{
	//				"type": "text",
	//			},
	//			"content": map[string]interface{}{
	//				"type": "text",
	//			},
	//			"location": map[string]interface{}{
	//				"type": "geo_point",
	//			},
	//			"time": map[string]interface{}{
	//				"type": "long",
	//			},
	//			"date": map[string]interface{}{
	//				"type": "date",
	//			},
	//			"age": map[string]interface{}{
	//				"type": "keyword",
	//			},
	//		},
	//	},
	//}
	//url := index
	//注意 映射信息不能更新
	//es.SearchHttp("PUT", url, query)

    //添加或修改文档，没有索引创建
	//doc := map[string]interface{}{
	//	"title":    "你好",
	//	"content":  "中国美丽的城市",
	//	"location": "41.015, -75.011",
	//	"time":     time.Now().Unix(),
	//	"date":     time.Now(),
	//	"age":      20,
	//}

	//es.Index(index, "", doc)
	//es.IndexEspi(index, "idx5", doc)
	//删除索引
	//es.Delete(index, "idx3")
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "vvvvv我爱你！!!!",
	//		},
	//	},
	//}
	//indexArr := []string{index}
	//es.DeleteByQuery(indexArr, query)
	////搜索单个字段
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "我爱你中国",
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}

	//搜索多个字段
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"multi_match": map[string]interface{}{
	//			"query":  "中国",
	//			"fields": []string{"title", "content"},
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}
	//提高某个字段权重，可以使用 ^ 字符语法为单个字段提升权重，在字段名称的末尾添加 ^boost ，其中 boost 是一个浮点数：
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"multi_match": map[string]interface{}{
	//			"query":  "中国",
	//			"fields": []string{"title", "content^2"},
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}

	//显示所有的
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match_all": map[string]interface{}{},
	//	},
	//}
	//es.Search(index, query)
	//地理距离过滤器（ geo_distance ）以给定位置为圆心画一个圆，来找出那些地理坐标落在其中的文档：
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"bool": map[string]interface{}{
	//			"must": map[string]interface{}{
	//				"match_all": map[string]interface{}{}, // 这里不设置其他查询条件，所以匹配全部文档
	//			},
	//			"filter": map[string]interface{}{
	//				"geo_distance": map[string]interface{}{
	//					"distance": "100km",
	//					"location": map[string]interface{}{
	//						"lat": 40.715,
	//						"lon": -73.988,
	//					},
	//				}},
	//		},
	//	},
	//	"sort": map[string]interface{}{ // 设置排序条件
	//		"_geo_distance": map[string]interface{}{ //_geo_distance代表根据距离排序
	//			"location": map[string]interface{}{ //根据location存储的经纬度计算距离。
	//				"lat": 40.715,  //当前纬度
	//				"lon": -73.988, //当前经度
	//			},
	//			"order": "asc", // asc 表示升序，desc 表示降序
	//		}},
	//}
	//es.Search(index, query)
	//精确值查询
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"age": "20",
	//		},
	//	},

	//}
	//es.Search(index, query)
	//范围查询
	//通过range实现范围查询，类似SQL语句中的>, >=, <, <=表达式。
	//gte范围参数 - 等价于>=
	//lte范围参数 - 等价于 <=
	//范围参数可以只写一个，例如：仅保留 "gte": 10， 则代表 FIELD字段 >= 10
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"range": map[string]interface{}{
	//			"age": map[string]interface{}{
	//				"gte": "19",
	//				"lte": "20",
	//			},
	//		},
	//	},
	//}
	//es.Search(index, query)
	//组合查询 如果需要编写类似SQL的Where语句，组合多个字段的查询条件，可以使用bool语句。
	//"must": [], // must条件，类似SQL中的and, 代表必须匹配条件
	//"must_not": [], // must_not条件，跟must相反，必须不匹配条件
	//"should": [] // should条件，类似SQL中or, 代表匹配其中一个条件
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"age": "19",
						},
					},
					{
						"match": map[string]interface{}{
							"title": "中国",
						},
					},
				},
			},
		},
	}
	es.Search(index, query)

	//使用mysql的方式来请求
	//query := map[string]interface{}{
	//	"query": "select title from " + index + " where title like '%中国%'", //这里使用mysql的方式来请求，非常简单，符合开发习惯，简化es入门门槛，支持order，支持Limit，那么排序和分页就自己写好了
	//}
	//query := map[string]interface{}{
	//	"query": "select title from " + index + " where title = '中国'", //这里使用mysql的方式来请求，非常简单，符合开发习惯，简化es入门门槛，支持order，支持Limit，那么排序和分页就自己写好了
	//}
	//es.SearchEsapiSql(query)

	//测试分词
	//query := map[string]interface{}{
	//	"analyzer": "ik_smart", //智能分词用：ik_smart，最大化分词用：ik_max_word
	//	"text":     "中国华人民",
	//}
	//url := index + "_analyze?pretty=true"

	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "我爱你中国",
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}
	//url := index + "/_search"

	//es.SearchHttp("GET", url, query)

}

````



### 参考资料
https://github.com/elastic/go-elasticsearch
https://www.tizi365.com/archives/590.html
https://www.elastic.co/guide/cn/elasticsearch/guide/current/_creating_an_index.html

## links

- [目录](/zh/preface.md)
- 上一节：
- 下一节：

