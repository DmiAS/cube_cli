package iproto

import (
	"fmt"
	"net"

	"github.com/DmiAS/cube_cli/internal/app/models"
)

const (
	protoType       = "tcp"
	svcMsg    int32 = 0x00000001
)

type Response = interface{}

type Client struct {
	addr *net.TCPAddr
}

func NewClient(host, port string) (*Client, error) {
	strAddr := fmt.Sprintf("%s:%s", host, port)
	addr, err := net.ResolveTCPAddr(protoType, strAddr)
	if err != nil {
		return nil, err
	}
	return &Client{addr: addr}, nil
}

func (c *Client) Send(token, scope string) (Response, error) {
	conn, err := net.DialTCP(protoType, nil, c.addr)
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

func sendPacket(conn *net.TCPConn, token, scope string) error {

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

func packRequest(token, scope string) ([]byte, error) {
	binBody, err := packBody(token, scope)
	if err != nil {
		return nil, err
	}

	binHeader, err := packHeader(binBody)
	if err != nil {
		return nil, err
	}

	req := append(binHeader, binBody...)
	return req, nil
}

func packBody(token, scope string) ([]byte, error) {
	svcToken := strToProtoString(token)
	svcScope := strToProtoString(scope)

	return Marshal(&models.Request{
		SvcMsg: svcMsg,
		Token:  svcToken,
		Scope:  svcScope,
	})
}
