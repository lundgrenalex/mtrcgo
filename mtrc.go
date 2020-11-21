package main

import (
	"github.com/lundgrenalex/mtrcgo/api"
	"github.com/lundgrenalex/mtrcgo/storage"
	"log"
	"net/http"
)

func main() {
	var addr = "127.0.0.1:8080"

	crudStorage := storage.NewInMemMetricsStorage()
	log.Println("Starting app..")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.GetStatus(crudStorage, w, r)
	})
	http.HandleFunc("/handler/gauge", func(w http.ResponseWriter, r *http.Request) {
		api.StoreGauge(crudStorage, w, r)
	})
	log.Println("Server listening on: " + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
