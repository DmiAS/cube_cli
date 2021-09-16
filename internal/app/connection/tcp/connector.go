package tcp

import (
	"fmt"
	"net"

	"github.com/DmiAS/cube_cli/internal/app/connection"
)

type Connector struct {
	addr *net.TCPAddr
}

func NewConnector(host, port string) (*Connector, error) {
	strAddr := fmt.Sprintf("%s:%s", host, port)
	addr, err := net.ResolveTCPAddr(protoType, strAddr)
	if err != nil {
		return nil, err
	}
	return &Connector{addr: addr}, nil
}

func (c *Connector) Dial() (connection.Connection, error) {
	conn, err := net.DialTCP(protoType, nil, c.addr)
	if err != nil {
		return nil, err
	}

	return &Conn{conn: conn}, nil
}
