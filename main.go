package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("start httpd server failed, ERR: %s\n", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcon to my http_server \n"))
	for key, v := range r.Header {
		for _, header := range v {
			// fmt.Printf("Header key: %s , Header value: %s \n", key, v)
			w.Header().Set(key, header)
		}
	}
	os.Setenv("VERSION", "V1.0")
	version := os.Getenv("VERSION")
	w.Header().Add("VERSION", version)
	// fmt.Printf("os version:%s \n", version)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("my http_server is health! \n"))
	fmt.Fprintf(w, "my http_server is health!")
}
