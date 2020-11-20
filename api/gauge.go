package api

import (
	"github.com/lundgrenalex/mtrcgo/storage"
	"io/ioutil"
	"net/http"
)

func StoreGauge(crud storage.CRUDStorage, w http.ResponseWriter, r *http.Request) {
	// Detect http method
	if r.Method != http.MethodPost {
		SendResponse(HttpResponse{405, "Method Not Allowed!"}, w)
		return
	}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		SendResponse(HttpResponse{500, readErr.Error()}, w)
		return
	}

	metric, err := storage.NewGaugeMetric(body)
	if err != nil {
		SendResponse(HttpResponse{500, err.Error()}, w)
		return
	}

	// Validate
	validateError := metric.Validate()
	if validateError != nil {
		SendResponse(HttpResponse{400, validateError.Error()}, w)
		return
	}

	// Store
	crud.CreateRecord(metric)
	// Encode
	SendResponse(metric, w)
	return
}
