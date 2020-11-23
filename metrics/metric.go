package metrics

type SimpleMetric struct {
	Name        string            `json:"name,omitempty"`
	Date        int               `json:"date,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Value       float32           `json:"value,omitempty"`
	Description string            `json:"description,omitempty"`
}

type MetricsSlice []SimpleMetric

type Repository interface {
	Upsert(SimpleMetric) error
	Remove(SimpleMetric) error
	Dump() MetricsSlice
}
