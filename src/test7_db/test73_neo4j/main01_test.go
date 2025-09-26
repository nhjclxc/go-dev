package test73_neo4j

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"testing"
)

// go get github.com/neo4j/neo4j-go-driver/v5/neo4j

func TestMain01(t *testing.T) {
	// 连接配置
	uri := "neo4j://localhost:7687"
	username := "neo4j"
	password := "neo4j123"

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close(context.Background())

	// 测试连接
	session := driver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	result, err := session.Run(context.Background(), "RETURN 'Hello, Neo4j!' AS message", nil)
	if err != nil {
		log.Fatal(err)
	}

	for result.Next(context.Background()) {
		fmt.Println(result.Record().Values[0]) // 输出: Hello, Neo4j!
	}
}
