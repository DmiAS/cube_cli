package connection

type Connection interface {
	Write(data []byte) (int, error)
	Read() ([]byte, error)
	Close() error
	CloseWrite() error
}

type Connector interface {
	Dial() (Connection, error)
}
