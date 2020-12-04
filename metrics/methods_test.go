package metrics

import "testing"

func TestMetricExpose(t *testing.T) {

	metric := SimpleMetric{
		Name:  "test_metric",
		Value: 45,
		Labels: map[string]string{
			"rsc_metric": "3711",
		},
		Date: 1606907901,
	}

	want := "# TYPE test_metric \ntest_metric{rsc_metric=\"3711\"} 45.000000\n"
	res := MetricsSlice{metric}.Expose()

	if want != res {
		t.Errorf("Expose format was incorrect!")
	}

}

func TestMetricExposeWithEmptyLabels(t *testing.T) {

	metric := SimpleMetric{
		Type: "gauge",
		Name:  "test_metric",
		Value: 45,
		Date:  1606907901,
	}

	want := "# TYPE test_metric gauge\ntest_metric 45.000000\n"
	res := MetricsSlice{metric}.Expose()

	if want != res {
		t.Errorf("Expose format was incorrect!")
	}

}

func TestMetricMultipleLabels(t *testing.T) {

	metric := SimpleMetric{
		Type: "gauge",
		Name:  "test_metric",
		Value: 45,
		Date:  1606907901,
		Labels: map[string]string{
			"rsc_metric": "3711",
			"2nd_metric": "42",
		},
	}

	want := "# TYPE test_metric gauge\ntest_metric{rsc_metric=\"3711\",2nd_metric=\"42\"} 45.000000\n"
	res := MetricsSlice{metric}.Expose()

	if want != res {
		t.Errorf("Expose format was incorrect!")
	}

}


func TestMultipleMetrics(t *testing.T) {

	metric1 := SimpleMetric{
		Type: "gauge",
		Name:  "test_metric",
		Value: 45,
		Date:  1606907901,
		Labels: map[string]string{
			"rsc_metric": "3711",
		},
	}

	metric2 := SimpleMetric{
		Type: "gauge",
		Name:  "test_metric2",
		Value: 45,
		Date:  1606907901,
		Labels: map[string]string{
			"muchWowSuchLabel": "42",
		},
	}

	want := "# TYPE test_metric gauge\ntest_metric{rsc_metric=\"3711\"} 45.000000\n# TYPE test_metric2 gauge\ntest_metric2{muchWowSuchLabel=\"42\"} 45.000000\n"
	res := MetricsSlice{metric1, metric2}.Expose()

	if want != res {
		t.Errorf("Expose format was incorrect!")
	}

}

func TestValidate(t *testing.T) {

	// Empty field
	metricWithEmptyName := SimpleMetric{
		Type: "gauge",
		Name:  "",
		Value: 45,
		Date:  1606907901,
	}

	err := metricWithEmptyName.Validate()
	if err == nil {
		t.Errorf("Bad parser!")
	}

	// WrongName
	metricWithWrongName := SimpleMetric{
		Type: "gauge",
		Name:  "bad-metric-name",
		Value: 45,
		Date:  1606907901,
	}

	err = metricWithWrongName.Validate()
	if err == nil {
		t.Errorf("Bad parser!")
	}

	// Wrong labels
	metricWithwrongLabel := SimpleMetric{
		Type: "gauge",
		Name:  "test_metric",
		Value: 45,
		Labels: map[string]string{
			"rsc-metric": "3711",
		},
		Date: 1606907901,
	}

	err = metricWithwrongLabel.Validate()
	if err == nil {
		t.Errorf("Bad parser!")
	}

	// All it's ok
	metric := SimpleMetric{
		Type: "gauge",
		Name:  "test_metric",
		Value: 45,
		Labels: map[string]string{
			"rsc_metric": "3711",
		},
		Date: 1606907901,
	}

	err = metric.Validate()
	if err != nil {
		t.Errorf("Bad parser!")
	}

}

func TestHash(t *testing.T) {

	metric := SimpleMetric{
		Name:  "test_metric",
		Value: 45,
		Labels: map[string]string{
			"rsc_metric": "3711",
		},
		Date: 1606907901,
	}

	hashstring := metric.Hash()
    want := "271d9d1c422af447fcecc8a2cabfacf5290011a7343e1d49f727ae5853fae1a9"
    if (hashstring != want) {
        t.Errorf("Bad hash!")
    }

}
