package storage

import (
	"sync"

	"github.com/lundgrenalex/mtrcgo/metrics"
)

// MetricsMemoryStorage is a shared struct
type MetricsMemoryStorage struct {
	mx      sync.RWMutex
	metrics map[string]metrics.SimpleMetric
}

// Init is a method that provides access to shared store
func Init() metrics.Repository {
	return &MetricsMemoryStorage{
		metrics: make(map[string]metrics.SimpleMetric, 0),
	}
}

func (s *MetricsMemoryStorage) exists(m metrics.SimpleMetric) bool {
	_, ok := s.metrics[m.Hash()]
	return ok
}

// Upsert metric into storage
func (s *MetricsMemoryStorage) Upsert(m metrics.SimpleMetric) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	mh := m.Hash()

	sm, exists := s.metrics[mh]
	if !exists {
		s.metrics[mh] = m
		return nil
	}

	switch mt := sm.Type; mt {
	case "gauge":
		sm = m
	case "counter":
		m.Value += sm.Value
	}

	s.metrics[mh] = m
	return nil

}

// Remove metric from storage
func (s *MetricsMemoryStorage) Remove(m metrics.SimpleMetric) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if !s.exists(m) {
		return nil
	}
	delete(s.metrics, m.Hash())
	return nil
}

// Dump metrics from storage
func (s *MetricsMemoryStorage) Dump() metrics.Slice {
	// НЕ ГАРАНТИРУЕТ порядка в metrics.Slice
	s.mx.RLock()
	defer s.mx.RUnlock()
	totalMetrics := len(s.metrics)
	if totalMetrics == 0 {
		// Dummy empty array
		return make(metrics.Slice, 0)
	}
	metricsToReturn := make(metrics.Slice, totalMetrics)
	idx := 0
	for _, v := range s.metrics {
		metricsToReturn[idx] = v
		idx++
	}
	return metricsToReturn
}
