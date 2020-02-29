package graphite

import (
	"net"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: 12,
		Timeout:  10 * time.Second,
	}

	server, err := New(&conf)
	if err != ErrUnsupportedProto {
		t.Fail()
	}
	if server != nil {
		t.Fail()
	}
}

func TestNew_ProtocolTCP(t *testing.T) {
	var err error
	var tcpServer Graphite
	var l net.Listener

	conf := &Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	tcpServer, err = New(conf)
	if err == nil {
		t.Fail()
	}

	l, err = net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish
	defer func() {
		if err := l.Close(); err != nil {
			t.Log(err)
			t.Fail()
		}
	}()

	tcpServer, err = New(conf)
	if err != nil {
		t.Fail()
	}
	srv, ok := tcpServer.(*GraphiteTCP)

	if !ok {
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

func TestNew_ProtocolUDP(t *testing.T) {
	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolUDP,
	}

	udpServer, err := New(&conf)
	if err != nil {
		t.Fail()
	}

	srv, ok := udpServer.(*GraphiteUDP)

	if !ok {
		t.Fail()
	}
	if srv.prefix != conf.Prefix {
		t.Fail()
	}

	if srv.address != conf.Address {
		t.Fail()
	}
}

func TestNew_ProtocolStdout(t *testing.T) {
	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolStdout,
		Timeout:  10 * time.Second,
	}

	stdoutServer, err := New(&conf)
	if err != nil {
		t.Fail()
	}
	srv, ok := stdoutServer.(*GraphiteStdout)

	if !ok {
		t.Fail()
	}
	if srv.prefix != conf.Prefix {
		t.Fail()
	}

}
