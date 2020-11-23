package metrics

import (
	"errors"
	"regexp"
	"crypto/md5"
)

func (m *SimpleMetric) Hash() string {
	var text string;
	text += m.Name
	byte_text = []byte(text)
	hash = md5.Sum(data).(string)
    return hash
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
