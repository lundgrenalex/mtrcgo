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

func (m *SimpleMetric) Hash() string {
	var text string
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

func (m *SimpleMetric) Validate() error {

	// Name
	if len(m.Name) == 0 {
		return errors.New("Empty Name field!")
	}

	if !(nameValidation.Match([]byte(m.Name))) {
		return errors.New("Incorrect Name field: " + m.Name)
	}

	// Labels
	for k, v := range m.Labels {
		if !(nameValidation.Match([]byte(k))) {
			return errors.New("Incorrect field name in labels: " + k + ":" + v)
		}
	}

	return nil

}

// https://prometheus.io/docs/instrumenting/exposition_formats/
func Expose(m MetricsSlice) string {
	getLabels := func (l map[string]string) string {
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
		exposedMetrics += fmt.Sprintf("%s%s %f\n", v.Name, getLabels(v.Labels), v.Value)
	}
	return exposedMetrics
}
