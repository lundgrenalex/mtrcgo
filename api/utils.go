package api

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e HttpResponse) Response() (int, interface{}) {
	return e.Status, e
}

type HTTPResponsive interface {
	Response() (int, interface{})
}

func SendResponse(r HTTPResponsive, w http.ResponseWriter) {
	status, response := r.Response()
	message, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write([]byte(string(message)))
}
