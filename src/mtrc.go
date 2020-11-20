package main

import (
	"api"
	"log"
	"net/http"
)


func main() {
	log.Println("Starting app..")
	http.HandleFunc("/", api.GetStatus)
	http.HandleFunc("/handler/gauge", api.StoreGauge)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
