package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/test", handleRequest)
	http.HandleFunc("/health", handleHealth)

	port := ":8080"
	log.Printf("Backend server starting on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 获取 Trace ID
	traceID := r.Header.Get("X-Trace-Id")
	retryAfter := r.Header.Get("Retry-After")
	requestID := r.Header.Get("X-Request-ID")
	cacheControl := r.Header.Get("Cache-Control")
	CacheRule := r.Header.Get("X-Cache-Rule")
	CDNGateway := r.Header.Get("X-CDN-Gateway")

	// 记录请求日志
	log.Printf("[%s] Retry-After=[%s], X-Request-ID=[%s], "+
		"Cache-Control=[%s], X-Cache-Rule=[%s], X-CDN-Gateway=[%s] "+
		"%s %s from %s",
		traceID, retryAfter, requestID,
		cacheControl, CacheRule, CDNGateway,
		r.Method, r.URL.Path, r.RemoteAddr)

	// 返回响应
	response := map[string]interface{}{
		"trace_id":      traceID,
		"Retry-After":   retryAfter,
		"X-Request-ID":  requestID,
		"Cache-Control": cacheControl,
		"X-Cache-Rule":  CacheRule,
		"X-CDN-Gateway": CDNGateway,
		"path":          r.URL.Path,
		"method":        r.Method,
		"timestamp":     time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Backend-Server", "go-cache-service")
	json.NewEncoder(w).Encode(response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
