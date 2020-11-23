package metrics

import (
	"errors"
	"regexp"
	"crypto/sha256"
	"fmt"
	"sort"
)

func (m *SimpleMetric) Hash() string {
	var text string;
	text += m.Name

	// ASC Sorting json keys
	keys := make([]string, 0, len(m.Labels))
	for k := range m.Labels {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Making str sum
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

	nameRegex := regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
	if !(nameRegex.Match([]byte(m.Name))) {
		return errors.New("Incorrect Name field: " + m.Name)
	}

	// Labels
	for k, v := range m.Labels {
		if !(nameRegex.Match([]byte(k))) {
			return errors.New("Incorrect field name in labels: " + k + ":" + v)
		}
	}

	return nil

}
