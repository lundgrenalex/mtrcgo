package api

import "regexp"
import "fmt"

type metric interface {
	Validate() (string, bool)
}

func Validate(m metric) (string, bool) {
	message, ok := m.Validate()
	return message, ok
}

func (m Gauge) Validate() (string, bool) {

	// Name
	if len(m.Name) == 0 {
		return "Empty Name field!", false
	}

	nameRegex := regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
	if !(nameRegex.Match([]byte(m.Name))) {
		return "Incorrect Name field: " + m.Name, false
	}

	// Labels
	for k, v := range m.Labels {
		if !(nameRegex.Match([]byte(k))) {
			return "Incorrect field name in labels: " + k, false
		}
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	return "Validated!", true

}
