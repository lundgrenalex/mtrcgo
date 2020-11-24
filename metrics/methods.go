package metrics

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"sort"
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
