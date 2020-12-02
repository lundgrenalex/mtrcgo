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

func TestValidate(t *testing.T) {

    // Empty field
    metricWithEmptyName := SimpleMetric{
		Name:  "",
		Value: 45,
		Date:  1606907901,
	}

    err := metricWithEmptyName.Validate()
    if (err == nil) {
        t.Errorf("Bad parser!")
    }

    // WrongName
    metricWithWrongName := SimpleMetric{
		Name:  "bad-metric-name",
		Value: 45,
		Date:  1606907901,
	}

    err = metricWithWrongName.Validate()
    if (err == nil) {
        t.Errorf("Bad parser!")
    }

    // Wrong labels
    metricWithwrongLabel := SimpleMetric{
		Name:  "test_metric",
		Value: 45,
        Labels: map[string]string{
            "rsc-metric": "3711",
        },
		Date:  1606907901,
	}

    err = metricWithwrongLabel.Validate()
    if (err == nil) {
        t.Errorf("Bad parser!")
    }

    // All it's ok
    metric := SimpleMetric{
		Name:  "test_metric",
		Value: 45,
        Labels: map[string]string{
            "rsc_metric": "3711",
        },
		Date:  1606907901,
	}

    err = metric.Validate()
    if (err != nil) {
        t.Errorf("Bad parser!")
    }

}
