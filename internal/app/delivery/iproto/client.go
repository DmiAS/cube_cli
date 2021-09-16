package iproto

import "github.com/DmiAS/cube_cli/internal/app/connection"

const (
	svcMsg int32 = 0x00000001
)

type Response = interface{}

type Client struct {
	connector connection.Connector
}

func NewClient(connector connection.Connector) *Client {
	return &Client{connector: connector}
}

func (c *Client) Send(token, scope string) (Response, error) {
	conn, err := c.connector.Dial()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err := sendPacket(conn, token, scope); err != nil {
		return nil, err
	}

	resp, err := getResp(conn)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendPacket(conn connection.Connection, token, scope string) error {

	packet, err := packRequest(token, scope)
	if err != nil {
		return err
	}

	if _, err := conn.Write(packet); err != nil {
		return err
	}

	if err := conn.CloseWrite(); err != nil {
		return err
	}

	return nil
}
