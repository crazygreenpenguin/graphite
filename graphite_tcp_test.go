package graphite

import (
	"net"
	"testing"
	"time"
)

func TestNewGraphiteTCP(t *testing.T) {
	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	srv, err := NewGraphiteTCP(&conf)
	if err == nil {
		t.Fail()
	}

	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish
	defer l.Close()

	srv, err = NewGraphiteTCP(&conf)
	if err != nil {
		t.Fail()
	}

	if srv.prefix != conf.Prefix {
		t.Fail()
	}
	if srv.timeout != conf.Timeout {
		t.Fail()
	}
	if srv.address != conf.Address {
		t.Fail()
	}
}

func TestGraphiteTCP_Connect(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish

	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	srv, err := NewGraphiteTCP(&conf)
	if err != nil {
		t.Fail()
	}

	err = srv.Connect()
	if err != nil {
		t.Fail()
	}
	conf = Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
	}

	srv, err = NewGraphiteTCP(&conf)
	if err != nil {
		t.Fail()
	}

	err = srv.Connect()
	if err != nil {
		t.Fail()
	}
	if srv.timeout != defaultTimeout*time.Second {
		t.Fail()
	}

	err = l.Close()
	if err != nil {
		t.Fail()
	}

	if err = srv.Connect(); err == nil {
		t.Fail()
	}
}

func TestGraphiteTCP_Disconnect(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish
	defer l.Close()

	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	srv, err := NewGraphiteTCP(&conf)
	if err != nil {
		t.Error(err)
	}
	err = srv.Disconnect()
	if err != nil {
		t.Error(err)
	}
	if srv.conn != nil {
		t.Fail()
	}
}

func TestGraphiteTCP_SendMetric(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish

	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	srv, err := NewGraphiteTCP(&conf)
	if err != nil {
		t.Error(err)
	}

	metric := Metric{}
	err = srv.SendMetric(&metric)
	if err != nil {
		t.Error(err)
	}

	metric = Metric{
		Name:      "test1",
		Value:     34,
		Timestamp: 0,
	}
	err = srv.SendMetric(&metric)
	if err != nil {
		t.Error(err)
	}

	metric = Metric{
		Name:      "",
		Value:     34,
		Timestamp: time.Now().Unix(),
	}
	err = srv.SendMetric(&metric)
	if err != nil {
		t.Error(err)
	}

	metric = Metric{
		Name:      "test1",
		Value:     34,
		Timestamp: time.Now().Unix(),
	}

	err = srv.SendMetric(&metric)
	if err != nil {
		t.Error(err)
	}

	srv.prefix = ""

	l.Close()

	err = srv.SendMetric(&metric)
	if err == nil {
		t.Fail()
	}

}

func TestGraphiteStdout_SendMetrics(t *testing.T) {
	metrics := []Metric{
		{
			Name:      "test1",
			Value:     34,
			Timestamp: time.Now().Unix(),
		},
		{
			Name:      "test1",
			Value:     34,
			Timestamp: time.Now().Unix(),
		},
		{
			Name:      "test1",
			Value:     34,
			Timestamp: time.Now().Unix(),
		},
	}

	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish

	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	srv, err := NewGraphiteTCP(&conf)
	if err != nil {
		t.Error(err)
	}

	if err = srv.SendMetrics(&metrics); err != nil {
		t.Error(err)
	}

	l.Close()

	if err = srv.SendMetrics(&metrics); err == nil {
		t.Fail()
	}
}

func TestGraphiteTCP_SimpleSend(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish

	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	srv, err := NewGraphiteTCP(&conf)
	if err != nil {
		t.Error(err)
	}

	if err = srv.SimpleSend("test2", 12); err != nil {
		t.Error(err)
	}

	l.Close()

	if err = srv.SimpleSend("test2", 12); err == nil {
		t.Fail()
	}
}
