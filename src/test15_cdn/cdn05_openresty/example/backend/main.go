package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	cacheID = getEnv("CACHE_ID", "origin")
	port    = getEnv("PORT", "8086")

	cachePath = make(map[string]string)
	mu        sync.Mutex
	logger    *log.Logger
)

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func lowerHeaders(h http.Header) map[string]string {
	res := map[string]string{}
	for k, v := range h {
		res[strings.ToLower(k)] = strings.Join(v, ",")
	}
	return res
}

func handler(w http.ResponseWriter, r *http.Request) {

	parsed, _ := url.Parse(r.RequestURI)
	params := parsed.Query()

	status, _ := strconv.Atoi(getFirst(params, "status", "200"))
	cacheControl := getFirst(params, "cache_control", "")
	expires := getFirst(params, "expires", "")

	bodyObj := map[string]interface{}{
		"path":    parsed.Path,
		"query":   params,
		"headers": lowerHeaders(r.Header),
	}

	body, _ := json.Marshal(bodyObj)

	// cache logic
	cacheStatus := "HIT"

	mu.Lock()
	if _, ok := cachePath[parsed.Path]; !ok {
		cacheStatus = "MISS"
		cachePath[parsed.Path] = parsed.Path
	}
	mu.Unlock()

	// headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.Header().Set("X-Mock-Backend", "true")
	w.Header().Set("X-Cache-Status", cacheStatus)
	w.Header().Set("X-Real-IP", r.Header.Get("X-Real-IP"))

	if cacheControl != "" {
		w.Header().Set("Cache-Control", cacheControl)
	}

	if expires != "" {
		w.Header().Set("Expires", expires)
	}

	w.WriteHeader(status)

	logger.Printf("X-Cache-Status %s", cacheStatus)
	logger.Printf("打印所有响应头 X-Real-IP: %s", r.Header.Get("X-Real-IP"))

	logger.Println("打印所有请求 Headers:")
	for k, v := range r.Header {
		logger.Printf("\t%s -->>> %s", k, strings.Join(v, ","))
	}

	// MISS 回源
	if cacheStatus == "MISS" {
		reqOrigin(r.Header)
	}

	if r.Method != "HEAD" {
		w.Write(body)
	}

	logger.Println("\n\n")
}

func getFirst(params url.Values, key, def string) string {
	if v := params.Get(key); v != "" {
		return v
	}
	return def
}

func reqOrigin(headers http.Header) {

	logger.Println("开始回源请求", headers.Get("X-Request-ID"))
	logger.Println("结束回源请求", headers.Get("X-Request-ID"))

}

func main() {

	os.MkdirAll("./logs", os.ModePerm)

	logFile, err := os.OpenFile(
		"./logs/"+cacheID+".log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		panic(err)
	}

	logger = log.New(logFile, "", log.LstdFlags)

	http.HandleFunc("/", handler)

	logger.Println("Server start at port", port)

	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		panic(err)
	}
}
