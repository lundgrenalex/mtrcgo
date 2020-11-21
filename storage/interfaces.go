package storage

type CRUDStorage interface {
	CreateRecord(g SingleMetric)
	DeleteRecord(g SingleMetric)
	UpdateRecord(g SingleMetric)
	GetAllRecords() MetricsSlice
}

func NewInMemMetricsStorage() CRUDStorage {
	return &InMemMetricsStorage{
		metrics: make(map[string]SingleMetric, 0),
	}
}

type SingleMetric interface {
	Hash() string
	Marshall()	([]byte, error)
	Validate() error
	Response() (int, interface{})
}

func NewGaugeMetric(data []byte) (SingleMetric, error) {
	metric, err := UnmarshallGauge(data)
	if err != nil {
		return nil, err
	}
	return metric, nil
}



