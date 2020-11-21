package storage

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
)

type MetricsSlice []SingleMetric

func (g *MetricsSlice) Marshall() []byte {
	msgBytes, err := json.Marshal(g)
	if err != nil {
		log.Fatal(err)
	}
	return msgBytes
}

type Gauge struct {
	Name   string            `json:"name,omitempty"`
	Date   int               `json:"date,omitempty"`
	Value  float32           `json:"value,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
}

func UnmarshallGauge(rawData []byte) (*Gauge, error) {
	msg := &Gauge{}
	err := json.Unmarshal(rawData, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (g *Gauge) Marshall() ([]byte, error) {
	msgBytes, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return msgBytes, nil
}

func (g Gauge) Response() (int, interface{}) {
	return 200, g
}

func (g *Gauge) Hash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", g)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (g *Gauge) Validate() error {

	// Name
	if len(g.Name) == 0 {
		return errors.New("Empty Name field!")
	}

	nameRegex := regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)
	if !(nameRegex.Match([]byte(g.Name))) {
		return errors.New("Incorrect Name field: " + g.Name)
	}

	// Labels
	for k, v := range g.Labels {
		if !(nameRegex.Match([]byte(k))) {
			return errors.New("Incorrect field name in labels: " + k + ":" + v)
		}
	}

	return nil

}
