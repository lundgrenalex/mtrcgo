package metrics

import "testing"

func TestMetricExpose(t *testing.T) {

	metric := SimpleMetric{
		Name:  "test_metric",
		Value: 45,
        Labels: map[string]string{
            "rsc_metric": "3711",
        },
		Date:  1606907901,
	}

	want := "test_metric{rsc_metric=\"3711\"} 45.000000\n"
	res := Expose(MetricsSlice{metric})

	if want != res {
		t.Errorf("Expose format was incorrect!")
	}

}

func TestMetricExposeWithEmptyLabels(t *testing.T) {

    metric := SimpleMetric{
		Name:  "test_metric",
		Value: 45,
		Date:  1606907901,
	}

	want := "test_metric 45.000000\n"
	res := Expose(MetricsSlice{metric})

	if want != res {
		t.Errorf("Expose format was incorrect!")
	}

}
