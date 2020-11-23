package main

import (
	"github.com/lundgrenalex/mtrcgo/api"
	"github.com/lundgrenalex/mtrcgo/storage"
	"log"
	"net/http"
)

const addr = "127.0.0.1:8080"

func main() {

	s := storage.Init()

	// Metrics
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.StoreMetric(s, w, r)
	})

	// Prometheus are watching only /metrics URL

	// Start webserver in goroutine
	log.Println("Server listening on: " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
