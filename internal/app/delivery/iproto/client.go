package iproto

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/DmiAS/cube_cli/internal/app/models"
)

const protoType = "tcp"

type Client struct {
	conn net.Conn
}

func NewClient(host, port string) (*Client, error) {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial(protoType, addr)
	if err != nil {
		return nil, err
	}
	return &Client{conn: conn}, nil
}

func (c *Client) Send(token, scope string) (*models.Response, error) {
	packet, err := packRequest(token, scope)
	if err != nil {
		return nil, err
	}

	_, err = c.conn.Write(packet)
	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(c.conn)
	if err != nil {
		return nil, err
	}

	var resPacket models.ResponsePacket
	if err := UnMarshal(resp, resPacket); err != nil {
		return nil, err
	}

	return &resPacket.Resp, nil
}

func (c *Client) Close() {
	c.conn.Close()
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

	req := append(binBody, binHeader...)
	return req, nil
}

func packBody(token, scope string) ([]byte, error) {
	svcToken := strToProtoString(token)
	svcScope := strToProtoString(scope)
	var svcMsg int32 = 0

	return Marshal(models.Request{
		SvcMsg: svcMsg,
		Token:  svcToken,
		Scope:  svcScope,
	})
}
