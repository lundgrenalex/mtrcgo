package storage

import (
    "sync"
	"github.com/lundgrenalex/mtrcgo/metrics"
)

type MetricsMemoryStorage struct {
	mx      sync.RWMutex
	metrics map[string]metrics.SimpleMetric
}

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

func (s *MetricsMemoryStorage) Upsert(m metrics.SimpleMetric) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.exists(m) {
		// Must update
		return nil
	}
	key := m.Hash()
	s.metrics[key] = m
	return nil
}

func (s *MetricsMemoryStorage) Remove(m metrics.SimpleMetric) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	return nil
}

func (s *MetricsMemoryStorage) Dump() metrics.MetricsSlice {
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
		idx += 1
	}
	return metricsToReturn
}
