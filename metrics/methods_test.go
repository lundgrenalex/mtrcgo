package metrics

import "testing"

func HashTest(t *testing.T) {
	metric := Metric{"hello_world", 1231434354, "", 1}
	hash := metric.hash()
	if hash == "" {
		t.Errorf("Abs(-1) = %d; want 1", hash)
	}
}
