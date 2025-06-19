package main

import (
	"encoding/json"
	"net/http"
)

// 获取HTTP请求的IP地址
func main03() {

	// 这篇文章演示了如何在Go中获取传入HTTP请求的IP地址。作为一种功能，它尝试使用X-FORWARDED-FORhttp头作为代理和负载均衡器（例如在Heroku之类的主机上）后面的代码，而RemoteAddr如果找不到该头，则会尝试使用http头。
	//
	//举个例子，我们在下面创建了一个（各种各样的）回显服务器，以json形式使用请求的IP地址回复传入的请求。

	http.HandleFunc("/", ExampleHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": GetIP(r),
	})
	w.Write(resp)
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
