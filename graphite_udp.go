package graphite

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// NewGraphiteUDP is a factory method that's used to create a new GraphiteUDP struct

func NewGraphiteUDP(conf *Config) (*GraphiteUDP, error) {
	server := GraphiteUDP{
		address: conf.Address,
		prefix:  conf.Prefix,
	}
	err := server.Connect()
	if err != nil {
		return nil, err
	}
	return &server, nil
}

// GraphiteUDP is a struct that defines TCP graphite connection

type GraphiteUDP struct {
	address string
	prefix  string
	conn    net.Conn
	lock    sync.Mutex
}

// Connect UDP to metric server

func (graphite *GraphiteUDP) Connect() error {
	graphite.lock.Lock()
	defer graphite.lock.Unlock()

	if graphite.conn != nil {
		graphite.conn.Close()
	}

	udpAddr, err := net.ResolveUDPAddr("udp", graphite.address)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp",
		nil,
		udpAddr)
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		return err
	}

	graphite.conn = conn
	return nil
}

// Disconnect closes the GraphiteUDP.conn field

func (graphite *GraphiteUDP) Disconnect() error {
	graphite.lock.Lock()
	defer graphite.lock.Unlock()

	err := graphite.conn.Close()
	graphite.conn = nil
	return err
}

// SendMetric send one metric to metric server
func (graphite *GraphiteUDP) SendMetric(metric *Metric) error {
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

	_, err := graphite.conn.Write(sendingMetric.ToByte())
	if err != nil {
		return err
	}
	return nil
}

// SendMetrics method sends the many metrics to metric server

func (graphite *GraphiteUDP) SendMetrics(metrics *[]Metric) error {
	for _, metric := range *metrics {
		err := graphite.SendMetric(&metric)
		if err != nil {
			return err
		}
	}
	return nil
}

// The SimpleSend method can be used to just pass a metric name and value and
// have it be sent to the GraphiteUDP host with the current timestamp

func (graphite *GraphiteUDP) SimpleSend(name string, value interface{}) error {
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
