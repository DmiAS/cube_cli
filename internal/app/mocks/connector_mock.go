package mocks

import "github.com/DmiAS/cube_cli/internal/app/connection"

type DialFn func() (connection.Connection, error)
type ErrFn func() error
type ReadFn func([]byte) ([]byte, error)
type WriteFn func([]byte) (int, error)

type Connector struct {
	dialFn DialFn
}

type Connection struct {
	CloseWriteFn ErrFn
	ReadFn       ReadFn
	WriteFn      WriteFn
	data         []byte
}

func NewConnector(dial DialFn) *Connector {
	return &Connector{dialFn: dial}
}

func (c *Connector) Dial() (connection.Connection, error) {
	return c.dialFn()
}

func (c *Connection) Close() error {
	return nil
}

func (c *Connection) CloseWrite() error {
	return c.CloseWriteFn()
}

func (c *Connection) Read() ([]byte, error) {
	if c.ReadFn != nil {
		return c.ReadFn(c.data)
	}
	return c.data, nil
}

func (c *Connection) Write(data []byte) (int, error) {
	if c.WriteFn != nil {
		return c.WriteFn(data)
	}
	c.data = data
	return len(c.data), nil
}
