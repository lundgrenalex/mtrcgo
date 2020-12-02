package main

import (
	"log"
	"net/http"

	"github.com/lundgrenalex/mtrcgo/api"
	"github.com/lundgrenalex/mtrcgo/storage"
)

const addr = "127.0.0.1:8080"

func main() {

	s := storage.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.StoreMetric(s, w, r)
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		api.ExposeMetrics(s, w, r)
	})

	log.Println("Server listening on: " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
