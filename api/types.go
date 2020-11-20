package api

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e HttpResponse) Response() (int, interface{}) {
	return e.Status, e
}