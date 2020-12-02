package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lundgrenalex/mtrcgo/metrics"
)

func StoreMetric(s metrics.Repository, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		SendResponse(HttpResponse{405, "Method Not Allowed!"}, w)
		return
	}

	var metric metrics.SimpleMetric
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		SendResponse(HttpResponse{500, readErr.Error()}, w)
		return
	}

	err := json.Unmarshal(body, &metric)
	if err != nil {
		SendResponse(HttpResponse{500, err.Error()}, w)
		return
	}

	err = s.Upsert(metric)
	if err != nil {
		SendResponse(HttpResponse{500, err.Error()}, w)
		return
	}

	SendResponse(HttpResponse{200, "Metric was updated!"}, w)
	return

}

func ExposeMetrics(s metrics.Repository, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		SendResponse(HttpResponse{405, "Method Not Allowed!"}, w)
	}

	m := s.Dump()
	w.WriteHeader(200)
	w.Write([]byte(metrics.Expose(m)))
	return

}
