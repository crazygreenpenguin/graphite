package graphite

import (
	"fmt"
	"sync"
	"time"
)

// NewGraphiteStdout is a factory method that's used to create a new GraphiteStdout struct
func NewGraphiteStdout(conf *Config) (*GraphiteStdout, error) {
	server := GraphiteStdout{
		prefix: conf.Prefix,
	}

	return &server, nil
}

// GraphiteStdout is a struct that write metric in stdout
type GraphiteStdout struct {
	prefix string
	lock   sync.Mutex
}

// Connect dummy method for Graphite interface implement's
func (graphite *GraphiteStdout) Connect() error {
	return nil
}

// Disconnect dummy method for Graphite interface implement's
func (graphite *GraphiteStdout) Disconnect() error {
	return nil
}

// SendMetric send one metric to stdout
func (graphite *GraphiteStdout) SendMetric(metric *Metric) error {
	graphite.lock.Lock()
	defer graphite.lock.Unlock()

	sendingMetric := Metric{}

	if metric.Name == "" {
		return nil
	}

	if metric.Timestamp == 0 {
		sendingMetric.Timestamp = time.Now().Unix()
	} else {
		sendingMetric.Timestamp = metric.Timestamp
	}

	if graphite.prefix == "" {
		sendingMetric.Name = metric.Name
	} else {
		sendingMetric.Name = fmt.Sprintf("%s.%s", graphite.prefix, metric.Name)
	}

	fmt.Printf("%s %s=%v",
		time.Unix(metric.Timestamp, 0).Format("2006-01-02 15:04:05"),
		sendingMetric.
			Name, sendingMetric.Value)
	return nil
}

// SendMetrics method sends the many metrics to metric server
func (graphite *GraphiteStdout) SendMetrics(metrics *[]Metric) error {
	for _, metric := range *metrics {
		err := graphite.SendMetric(&metric)
		if err != nil {
			return err
		}
	}
	return nil
}

// The SimpleSend method can be used to just pass a metric name and value and
// have it be sent to the GraphiteStdout host with the current timestamp
func (graphite *GraphiteStdout) SimpleSend(name string, value interface{}) error {
	var metricName string
	if graphite.prefix == "" {
		metricName = name
	} else {
		metricName = fmt.Sprintf("%s.%s", graphite.prefix, name)
	}
	metric := NewMetric(metricName, value, time.Now().Unix())
	err := graphite.SendMetric(&metric)
	if err != nil {
		return err
	}
	return nil
}
