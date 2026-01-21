package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	serviceName := os.Getenv("NAME")
	if port == "" {
		port = "8080"
	}

	// http://127.0.0.1:8080/**
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "serviceName": serviceName})
	})

	log.Printf("🚀 Mock Server on http://0.0.0.0:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
