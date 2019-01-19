package graphite

import (
	"fmt"
)

// Metric is a struct that defines graphite metric
type Metric struct {
	Name      string
	Value     interface{}
	Timestamp int64
}

func NewMetric(name string, value interface{}, timestamp int64) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
}

func (metric Metric) ToString() string {
	return fmt.Sprintf(
		"%s %v %d",
		metric.Name,
		metric.Value,
		metric.Timestamp,
	)
}

func (metric Metric) ToByte() []byte {
	return []byte(fmt.Sprintf(
		"%s %v %d",
		metric.Name,
		metric.Value,
		metric.Timestamp,
	))

}
