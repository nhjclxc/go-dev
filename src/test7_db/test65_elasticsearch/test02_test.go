package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
)

// index 相关操作
func Test0201(t *testing.T) {
	es := getEsClient()

	// 2. 创建索引（带 mapping）
	indexName := "my_index"
	mapping := `{
	  "settings": {
	    "number_of_shards": 1,
	    "number_of_replicas": 1
	  },
	  "mappings": {
	    "properties": {
	      "title": { "type": "text" },
	      "tags":  { "type": "keyword" },
	      "date":  { "type": "date" },
	      "views": { "type": "integer" }
	    }
	  }
	}`

	res, err := es.Indices.Create(indexName, es.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Create Index:", res.String())

	// 3. 插入文档
	doc := map[string]interface{}{
		"title": "Elasticsearch 入门",
		"tags":  []string{"search", "es"},
		"date":  "2025-09-29",
		"views": 100,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf("Error encoding doc: %s", err)
	}

	res, err = es.Index(indexName, &buf, es.Index.WithDocumentID("1"))
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Index Doc:", res.String())

	// 4. 获取文档
	res, err = es.Get(indexName, "1")
	if err != nil {
		log.Fatalf("Error getting document: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Get Doc:", res.String())

	// 5. 搜索文档
	query := `{
	  "query": {
	    "match": {
	      "title": "Elasticsearch"
	    }
	  }
	}`
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error searching: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Search Result:", res.String())

}
