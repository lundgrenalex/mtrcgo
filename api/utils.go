package api

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse - template for http response
type HTTPResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Response - interface for API error answers
func (e HTTPResponse) Response() (int, interface{}) {
	return e.Status, e
}

// Response is a interface for Response
type Response interface {
	Response() (int, interface{})
}

// SendResponse - output writer for webserver
func SendResponse(r HTTPResponse, w http.ResponseWriter) {
	status, response := r.Response()
	message, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write([]byte(string(message)))
}
