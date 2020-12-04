package metrics

// SimpleMetric is a basic metric struct
type SimpleMetric struct {
	Type        string            `json:"type,omitempty"`
	Name        string            `json:"name,omitempty"`
	Date        int               `json:"date,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Value       float32           `json:"value,omitempty"`
	Description string            `json:"description,omitempty"`
}

// MetricsSlice is a slice of SimpleMetric
type MetricsSlice []SimpleMetric

// Repository is a repository of methods for storing our metrics
type Repository interface {
	Upsert(SimpleMetric) error
	Remove(SimpleMetric) error
	Dump() MetricsSlice
}
