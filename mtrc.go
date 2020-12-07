package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lundgrenalex/mtrcgo/api"
	"github.com/lundgrenalex/mtrcgo/storage"
)

// Config struct
type Config struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port string `yaml:"port"`
	} `yaml:"server"`

	SnapShot struct{
		FilePath string `yaml:"file"`
		// Delta is a time between snapshots in seconds
		Delta	time.Duration `yaml:"delta"`
	} `yaml:"snapshots"`
}

func main() {

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	config, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	s := storage.Init()
	s.LoadSnapShot(config.SnapShot.FilePath)
	go s.DumpSnapShot(config.SnapShot.Delta, config.SnapShot.FilePath)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.StoreMetric(s, w, r)
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		api.ExposeMetrics(s, w, r)
	})

	addr := config.Server.Host + ":" + config.Server.Port

	log.Println("Server listening on: " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
