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

// NewMetric generate new Metric struct
func NewMetric(name string, value interface{}, timestamp int64) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
}

// ToString convert Metric struct to string view
func (metric Metric) ToString() string {
	return fmt.Sprintf(
		"%s %v %d",
		metric.Name,
		metric.Value,
		metric.Timestamp,
	)
}

// ToByte convert Metric struct to []byte ready for remote sending
func (metric Metric) ToByte() []byte {
	return []byte(fmt.Sprintf(
		"%s %v %d\n",
		metric.Name,
		metric.Value,
		metric.Timestamp,
	))

}
