package storage

import (
    "testing"
    "github.com/lundgrenalex/mtrcgo/metrics"
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
