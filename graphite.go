package graphite

import (
	"errors"
	"time"
)

const (
	ProtocolTCP = 0
	//ProtocolTCP define graphite server protocol TCP
	ProtocolUDP = 1
	//ProtocolUDP define graphite server protocol UDP
	ProtocolStdout = 2
	//ProtocolStdout define printing metric in stdout
)

//Error unsupported protocol

var ErrUnsupportedProto = errors.New("unsupported protocol")

//Interface to send metric to graphite server
type Graphite interface {
	// SendMetric send one metric to metric server
	SendMetric(metric *Metric) error
	// SendMetric send some metrics to metric server
	SendMetrics(metric *[]Metric) error
	// Connect to metric server
	Connect() error
	// SimpleSend generate metric with now Timestamp, and send this to metric server
	SimpleSend(name string, value interface{}) error
}

// Config configurable parameter of Graphite interface
type Config struct {
	// Address - address of metric sever in host:port format. For example 11.11.11.11:2003
	Address string
	// Protocol - set protocol will be using to metric send, graphite.ProtocolTCP for example.
	Protocol uint8
	// Timeout - set send timeout TCP
	Timeout time.Duration
	// Prefix - add prefix to metric.Name, if name metric is test1 and prefix test_metric,
	// then send metric with name  test_metric.test1
	Prefix string
}

//New return Graphite struct with config setting by conf

func New(conf *Config) (Graphite, error) {
	switch conf.Protocol {
	case ProtocolTCP:
		return NewGraphiteTCP(conf)
	case ProtocolUDP:
		return NewGraphiteUDP(conf)
	case ProtocolStdout:
		return NewGraphiteStdout(conf)
	default:
		return nil, ErrUnsupportedProto
	}
}
