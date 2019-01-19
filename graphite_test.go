package graphite

import (
	"net"
	"testing"
	"time"
)

func TestNew_ProtocolTCP(t *testing.T) {
	conf := Config{
		Address:  "127.0.0.1:3300",
		Prefix:   "test",
		Protocol: ProtocolTCP,
		Timeout:  10 * time.Second,
	}

	tcpServer, err := New(&conf)
	if err == nil {
		t.Fail()
	}

	l, err := net.Listen("tcp", "127.0.0.1:3300")
	if err != nil {
		t.Error(err)
	}
	// Close the listener when test finish
	defer l.Close()

	tcpServer, err = New(&conf)
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