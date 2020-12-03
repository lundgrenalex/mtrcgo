package storage

import (
	"sync"

	"github.com/lundgrenalex/mtrcgo/metrics"
)

// MetricsMemoryStorage is shared struct
type MetricsMemoryStorage struct {
	mx      sync.RWMutex
	metrics map[string]metrics.SimpleMetric
}

// Init is a method that provide access to shared store
func Init() metrics.Repository {
	return &MetricsMemoryStorage{
		metrics: make(map[string]metrics.SimpleMetric, 0),
	}
}

func (s *MetricsMemoryStorage) exists(m metrics.SimpleMetric) bool {
	if _, ok := s.metrics[m.Hash()]; ok {
		return true
	}
	return false
}

// Upsert is a method than store metric in wh
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
	case "guage":
		sm = m
	case "counter":
		m.Value += sm.Value
	}

	s.metrics[mh] = m
	return nil

}

// Remove is a method for removing metrics
func (s *MetricsMemoryStorage) Remove(m metrics.SimpleMetric) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if !s.exists(m) {
		return nil
	}
	delete(s.metrics, m.Hash())
	return nil
}

// Dump is method for dumping metrics from shared struct
func (s *MetricsMemoryStorage) Dump() metrics.MetricsSlice {
	// НЕ ГАРАНТИРУЕТ порядка в metrics.MetricsSlice
	s.mx.RLock()
	defer s.mx.RUnlock()
	totalMetrics := len(s.metrics)
	if totalMetrics == 0 {
		// Dummy empty array
		return make(metrics.MetricsSlice, 0)
	}
	metricsToReturn := make(metrics.MetricsSlice, totalMetrics)
	idx := 0
	for _, v := range s.metrics {
		metricsToReturn[idx] = v
		idx++
	}
	return metricsToReturn
}
