package tcp

import (
	"io/ioutil"
	"net"
)

const protoType = "tcp"

type Conn struct {
	conn *net.TCPConn
}

func (t *Conn) Write(data []byte) (int, error) {
	return t.conn.Write(data)
}

func (t *Conn) Read() ([]byte, error) {
	return ioutil.ReadAll(t.conn)
}

func (t *Conn) Close() error {
	return t.conn.Close()
}

func (t *Conn) CloseWrite() error {
	return t.conn.CloseWrite()
}
