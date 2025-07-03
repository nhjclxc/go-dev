package main

import "fmt"

func main() {
	// 示例数据
	data := []map[string]interface{}{
		{"id": "123", "name": "Alice"},
		{"id": "456", "name": "Bob"},
		{"id": "789", "name": "Charlie"},
	}

	var ids []interface{}
	for _, item := range data {
		if id, ok := item["id"]; ok {
			ids = append(ids, id)
		}
	}

	fmt.Println("IDs:", ids)
}
