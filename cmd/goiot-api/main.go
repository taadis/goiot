package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "index")
	})
	err := http.ListenAndServe(":8089", mux)
	if err != nil {
		log.Fatalf("http.ListenAndServe error:%+v", err)
	}
}
