package storage

import (
	"sync"
)


type InMemMetricsStorage struct {
	mx sync.RWMutex
	metrics map[string]SingleMetric

}

func (ms *InMemMetricsStorage) isGaugeExists(g SingleMetric) bool {
	for h, _ := range ms.metrics {
		if h == g.Hash(){
			return true
		}
	}
	return false
}

func (ms *InMemMetricsStorage) CreateRecord(g SingleMetric) {
	ms.mx.Lock()
	defer ms.mx.Unlock()
	if ms.isGaugeExists(g) {
		return
	}
	key := g.Hash()
	ms.metrics[key] = g
}

func (ms *InMemMetricsStorage) DeleteRecord(g SingleMetric) {
	ms.mx.Lock()
	defer ms.mx.Unlock()
	if !ms.isGaugeExists(g) {
		return
	}
	delete(ms.metrics, g.Hash())
}

func (ms *InMemMetricsStorage) UpdateRecord(g SingleMetric) {
	ms.mx.Lock()
	defer ms.mx.Unlock()
	if !ms.isGaugeExists(g) {
		return
	}
	ms.metrics[g.Hash()] = g
}

func (ms *InMemMetricsStorage) GetAllRecords() MetricsSlice {
	ms.mx.RLock()
	defer ms.mx.RUnlock()
	totalMetrics := len(ms.metrics)
	if totalMetrics == 0 {
		// Dummy empty array
		return make(MetricsSlice, 0)
	}
	metricsToReturn := make(MetricsSlice, totalMetrics)
	idx := 0
	for _, v := range ms.metrics {
		metricsToReturn[idx] = v
		idx +=1
	}
	return metricsToReturn
}

