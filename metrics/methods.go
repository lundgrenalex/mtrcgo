package metrics

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	// Metric Regexp Validation
	nameValidation = regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
)

// Hash is a hash function for metric
func (m *SimpleMetric) Hash() string {
	var text string
	text += m.Type
	text += m.Name
	keys := make([]string, 0, len(m.Labels))
	for k := range m.Labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		text += k + m.Labels[k]
	}

	h := sha256.New()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Validate will check metric
func (m *SimpleMetric) Validate() error {

	// Name
	if len(m.Name) == 0 {
		return errors.New("Empty name field")
	}

	if !(nameValidation.Match([]byte(m.Name))) {
		return errors.New("Incorrect Name field: " + m.Name)
	}

	// Types: Gauge, Counter or Histogram
	checkMetricType := func(t string) bool {
		switch t {
		case
			"gauge",
			"counter",
			"histogram":
			return true
		}
		return false

	}

	if !checkMetricType(m.Type) {
		return errors.New("Wrong type: " + m.Type + ". Type must be counter, histogram or guage")
	}

	// Labels
	for k, v := range m.Labels {
		if !(nameValidation.Match([]byte(k))) {
			return errors.New("Incorrect field name in labels: " + k + ":" + v)
		}
	}

	return nil

}

// Expose func for metrics
// https://prometheus.io/docs/instrumenting/exposition_formats/
func (m MetricsSlice) Expose() string {
	getLabels := func(l map[string]string) string {
		if l == nil {
			return ""
		}
		var labels = make([]string, len(l))
		i := 0
		for k, v := range l {
			labels[i] = fmt.Sprintf("%s=\"%s\"", k, v)
			i++
		}
		return fmt.Sprintf("{%s}", strings.Join(labels, ","))
	}

	var exposedMetrics string
	for _, v := range m {
		if v.Type == "counter" {
			v.Name += "_total"
		}
		if v.Description != "" {
			exposedMetrics += fmt.Sprintf("# HELP %s %s\n", v.Name, v.Description)
		}
		exposedMetrics += fmt.Sprintf("# TYPE %s %s\n", v.Name, v.Type)
		exposedMetrics += fmt.Sprintf("%s%s %f\n", v.Name, getLabels(v.Labels), v.Value)
	}
	return exposedMetrics
}
