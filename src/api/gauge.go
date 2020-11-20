package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func StoreGauge(w http.ResponseWriter, r *http.Request) {
	// Detect http method
	if r.Method != http.MethodPost {
		SendResponse(HttpResponse{405, "Method Not Allowed!"}, w)
		return
	}

	// Decode incoming data
	var metric Gauge

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		SendResponse(HttpResponse{500, readErr.Error()}, w)
		return
	}

	decodeErr := json.Unmarshal(body, &metric)
	if decodeErr != nil {
		SendResponse(HttpResponse{500, decodeErr.Error()}, w)
		return
	}

	// Validate
	validateError := Validate(metric)
	if validateError != nil {
		SendResponse(HttpResponse{400, validateError.Error()}, w)
		return
	}

	// Store

	// Encode
	SendResponse(metric, w)
	return

}
