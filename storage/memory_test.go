package storage

import (
	"fmt"
	"github.com/lundgrenalex/mtrcgo/metrics"
	"sync"
	"testing"
)

func TestStorage(t *testing.T) {

    s := Init()
    if s == nil {
        t.Errorf("I can't init storage!")
    }

    metric := metrics.SimpleMetric{
        Name:  "test_metric",
		Value: 45,
		Labels: map[string]string{
			"rsc_metric": "3711",
		},
		Date: 1606907901,
    }

    res := s.Upsert(metric)
    if res != nil {
		t.Errorf("I can't insert new metric to store!")
	}

    allMetrics := s.Dump()
    if allMetrics == nil {
		t.Errorf("I can't dump metric from store!")
	}

    res = s.Remove(metric)
    if res != nil {
		t.Errorf("I can't remove metric from store!")
	}

    // Try to remove if metric not exists
    res = s.Remove(metric)
    if res != nil {
		t.Errorf("Got error while we removed metric from store!")
	}

    // Dump empty store
    allMetrics = s.Dump()
    if allMetrics == nil {
		t.Errorf("I can't dump metric from store!")
	}

}

func TestConcurrentStorage(t *testing.T) {
	metricsToRemove := 499
	metricsToCreate := 5000
	var wg sync.WaitGroup

	s := Init()
	if s == nil {
		t.Errorf("I can't init storage!")
	}
	
	m := make(metrics.MetricsSlice, metricsToCreate)
	for i := 0; i < metricsToCreate; i++ {
		m[i] = metrics.SimpleMetric{
			Name:  "test_metric",
			Value: float32(i),
			Labels: map[string]string{
				fmt.Sprintf("%d", i): "3711",
			},
			Date: 1606907901,
		}
	}

	wg.Add(len(m))
	for _, mtrc := range m {
		go func(i metrics.SimpleMetric) {
			defer wg.Done()
			s.Upsert(i)
		}(mtrc)
	}
	wg.Wait()

	wg.Add(metricsToRemove)

	for i := 0; i < metricsToRemove; i ++ {
		go func(idx int) {
			defer wg.Done()
			s.Remove(m[idx+idx])
		}(i)
	}
	wg.Wait()
	if len(s.Dump()) != metricsToCreate - metricsToRemove {
		t.Errorf("Upserted metrics count != removed metrics count")
	}
}