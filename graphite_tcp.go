package graphite

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const defaultTimeout = 5

// defaultTimeout is the default number of seconds that we're willing to wait
// before forcing the connection establishment to fail

// NewGraphiteTCP is a factory method that's used to create a new GraphiteTCP struct
func NewGraphiteTCP(conf *Config) (*GraphiteTCP, error) {
	server := GraphiteTCP{
		address: conf.Address,
		prefix:  conf.Prefix,
		timeout: conf.Timeout,
	}
	err := server.Connect()
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// GraphiteTCP is a struct that defines TCP graphite connection
type GraphiteTCP struct {
	address string
	timeout time.Duration
	prefix  string
	conn    net.Conn
	lock    sync.Mutex
}

// Connect establish TCP connection to metric server
func (graphite *GraphiteTCP) Connect() error {
	graphite.lock.Lock()
	defer graphite.lock.Unlock()

	if graphite.conn != nil {
		if err := graphite.conn.Close(); err != nil {
			return err
		}
	}

	if graphite.timeout == 0 {
		graphite.timeout = defaultTimeout * time.Second
	}

	conn, err := net.DialTimeout(
		"tcp",
		graphite.address,
		graphite.timeout)
	if err != nil {
		return err
	}

	graphite.conn = conn
	return nil
}

// Disconnect closes the GraphiteTCP.conn field
func (graphite *GraphiteTCP) Disconnect() error {
	graphite.lock.Lock()
	defer graphite.lock.Unlock()

	err := graphite.conn.Close()
	graphite.conn = nil
	return err
}

// SendMetric send one metric to metric server
func (graphite *GraphiteTCP) SendMetric(metric *Metric) error {
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

	sendingMetric.Value = metric.Value

	_, err := graphite.conn.Write(sendingMetric.ToByte())
	if err != nil {
		return err
	}
	return nil
}

// SendMetrics method sends the many metrics to metric server
func (graphite *GraphiteTCP) SendMetrics(metrics *[]Metric) error {
	for _, metric := range *metrics {
		err := graphite.SendMetric(&metric)
		if err != nil {
			return err
		}
	}
	return nil
}

// SimpleSend method can be used to just pass a metric name and value and
// have it be sent to the GraphiteTCP host with the current timestamp
func (graphite *GraphiteTCP) SimpleSend(name string, value interface{}) error {
	metric := NewMetric(name, value, time.Now().Unix())
	err := graphite.SendMetric(&metric)
	if err != nil {
		return err
	}
	return nil
}
