package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lundgrenalex/mtrcgo/metrics"
)

// StoreMetric i a method for webserver
func StoreMetric(s metrics.Repository, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		SendResponse(HTTPResponse{405, "Method Not Allowed!"}, w)
		return
	}

	var metric metrics.SimpleMetric
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		SendResponse(HTTPResponse{500, readErr.Error()}, w)
		return
	}

	err := json.Unmarshal(body, &metric)
	if err != nil {
		SendResponse(HTTPResponse{500, err.Error()}, w)
		return
	}

	err = metric.Validate()
	if err != nil {
		SendResponse(HTTPResponse{500, err.Error()}, w)
		return
	}

	err = s.Upsert(metric)
	if err != nil {
		SendResponse(HTTPResponse{500, err.Error()}, w)
		return
	}

	SendResponse(HTTPResponse{200, "Metric was updated!"}, w)
	return

}

// ExposeMetrics is a interface for Prometheus scrapers
func ExposeMetrics(s metrics.Repository, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		SendResponse(HTTPResponse{405, "Method Not Allowed!"}, w)
	}

	m := s.Dump()
	w.WriteHeader(200)
	w.Write([]byte(m.Expose()))
	return

}
