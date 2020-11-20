package api

import "encoding/json"
import "net/http"
import "io"

type response interface {
	GetResponse() (int, interface{})
}

func (e HttpResponse) GetResponse() (int, interface{}) {
	return e.Status, e
}

func (m Gauge) GetResponse() (int, interface{}) {
	return 200, m
}

func SendResponse(r response, w http.ResponseWriter) {
	status, response := r.GetResponse()
	message, _ := json.Marshal(response)
	w.WriteHeader(status)
	w.Write([]byte(string(message)))
}

func DecodeJSON(s *io.ReadCloser) string {
	return ""
}
