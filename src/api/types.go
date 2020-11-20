package api

type Gauge struct {
	Name   string            `json:"name,omitempty"`
	Date   int               `json:"date,omitempty"`
	Value  float32           `json:"value,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
