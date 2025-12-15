package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// 不断从192.168.203.182:9092里面的 nginx_error 里面读取消息
func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"192.168.203.182:9092"},
		Topic:   "nginx_error",
	})

	defer reader.Close()

	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("read message error: %v", err)
		}

		fmt.Printf(
			"topic=%s partition=%d offset=%d value=%s\n",
			msg.Topic,
			msg.Partition,
			msg.Offset,
			string(msg.Value),
		)
	}
}

// {"@timestamp":"2025-12-15T05:53:32.383Z"}
// 将这个json消息追加写入/home/lxc/cdn/cdn-gateway/scripts/logs/pcdn-access-mock.log文件
// echo '{"@timestamp":"2025-12-15T13:55:32.383Z",name:"lxc"}' >> /home/lxc/cdn/cdn-gateway/scripts/logs/pcdn-access-mock.log
/*
topic=nginx_error partition=0 offset=28 value={"@timestamp":"2025-12-15T05:53:32.383Z","@metadata":{"beat":"filebeat","type":"_doc","version":"8.12.0"},"ecs":{"version":"8.0.0"},"host":{"name":"b3f341a479a8"},"log":{"file":{"path":"/var/log/nginx/pcdn-access-mock.log","device_id":"64514","inode":"4875800"},"offset":1543},"message":"{\"trace_id\": \"d11a1f8568f0afd5c27b1f75e2af0ff4\", \"time\": \"2025-12-13T05:53:29+08:00\", \"client_ip\": \"10.0.55.59\", \"x_forwarded_for\": \"\", \"host\": \"cdn.customer-e.com\", \"method\": \"GET\", \"uri\": \"/live/segment_66916.mp4\", \"args\": \"\", \"protocol\": \"HTTP/1.1\", \"scheme\": \"https\", \"range\": \"\", \"request_length\": 543, \"status\": 200, \"body_bytes_sent\": 30846218, \"request_time\": 2.722, \"upstream_addr\": \"127.0.0.1:8080\", \"upstream_status\": \"200\", \"upstream_response_time\": \"2.681\", \"upstream_connect_time\": \"0.041\", \"upstream_header_time\": \"0.027\", \"cache_status\": \"HIT\", \"connection\": \"97967\", \"connection_requests\": 2, \"user_agent\": \"Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36\", \"referer\": \"https://cdn.customer-e.com/player.html\"}","input":{"type":"filestream"},"agent":{"id":"1e18a76d-875b-4339-9e6a-222b96d3ce04","name":"b3f341a479a8","type":"filebeat","version":"8.12.0","ephemeral_id":"242d78d2-a723-4f1d-9e0c-a8ac7966c74c"}}
topic=nginx_error partition=0 offset=30 value={"@timestamp":"2025-12-15T05:56:20.412Z","@metadata":{"beat":"filebeat","type":"_doc","version":"8.12.0"},"ecs":{"version":"8.0.0"},"host":{"name":"b3f341a479a8"},"agent":{"id":"1e18a76d-875b-4339-9e6a-222b96d3ce04","name":"b3f341a479a8","type":"filebeat","version":"8.12.0","ephemeral_id":"242d78d2-a723-4f1d-9e0c-a8ac7966c74c"},"log":{"offset":2383,"file":{"path":"/var/log/nginx/pcdn-access-mock.log","device_id":"64514","inode":"4875800"}},"message":"{\"@timestamp\":\"2025-12-15T13:55:32.383Z\",name:\"lxc\"}","input":{"type":"filestream"}}
*/
