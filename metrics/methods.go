package metrics

import (
	"errors"
	"regexp"
	"crypto/sha256"
	"fmt"
)

func (m *SimpleMetric) Hash() string {
	var text string;
	text += m.Name
	for k, v := range m.Labels {
		text += k + v
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
