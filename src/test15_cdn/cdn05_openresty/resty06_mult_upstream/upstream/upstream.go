package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {

	// HOST=192.168.201.74 PORT=8080 NAME=upstream go run upstream.go
	// HOST=192.168.201.74 PORT=8081 NAME=upstream1 go run upstream.go

	hosttmp := os.Getenv("HOST")
	port := os.Getenv("PORT")
	serviceName := os.Getenv("NAME")

	host := "0.0.0.0"
	if hosttmp != "" {
		host = hosttmp
	}
	if port == "" {
		port = "8080"
	}

	// http://127.0.0.1:8080/**
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "serviceName": serviceName})
	})

	log.Printf("🚀 Mock Server on http://%s:%s", host, port)
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
