package api

import (
	"encoding/json"
	"net/http"
)

type HTTPResponsive interface {
	Response() (int, interface{})
}

func SendResponse(r HTTPResponsive, w http.ResponseWriter) {
	status, response := r.Response()
	message, _ := json.Marshal(response)
	w.WriteHeader(status)
	_, _ = w.Write([]byte(string(message)))
}
