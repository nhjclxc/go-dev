package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/health", handleHealth)

	port := ":8080"
	log.Printf("Backend server starting on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 获取 Trace ID
	traceID := r.Header.Get("X-Trace-Id")

	// 记录请求日志
	log.Printf("[%s] %s %s from %s", traceID, r.Method, r.URL.Path, r.RemoteAddr)

	// 返回响应
	response := map[string]interface{}{
		"message":   "Hello World from Backend Cache Service",
		"trace_id":  traceID,
		"path":      r.URL.Path,
		"method":    r.Method,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Backend-Server", "go-cache-service")
	json.NewEncoder(w).Encode(response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
