package iproto

import (
	"fmt"
	"io/ioutil"
	"net"

	"github.com/DmiAS/cube_cli/internal/app/models"
)

const (
	protoType       = "tcp"
	svcMsg    int32 = 0x00000001
)

type Client struct {
	addr *net.TCPAddr
}

type Response = interface{}

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

	packet, err := packRequest(token, scope)
	if err != nil {
		return nil, err
	}

	if err := sendPacket(conn, packet); err != nil {
		return nil, err
	}

	resp, err := getResp(conn)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func sendPacket(conn *net.TCPConn, packet []byte) error {
	if _, err := conn.Write(packet); err != nil {
		return err
	}

	if err := conn.CloseWrite(); err != nil {
		return err
	}

	return nil
}

func getResp(conn *net.TCPConn) (Response, error) {
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	resPacket := new(models.ResponsePacket)
	if err := UnMarshal(data, resPacket); err != nil {
		return nil, err
	}

	resp, err := determineResponse(resPacket.Body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func determineResponse(packet models.ResponseBody) (Response, error) {
	if packet.ReturnCode != 0 {
		errString, err := packet.ErrorString.ToString()
		if err != nil {
			return nil, err
		}
		return models.ResponseErr{
			ReturnCode:  packet.ReturnCode,
			ErrorString: errString,
		}, nil
	}

	errString, err := packet.ErrorString.ToString()
	if err != nil {
		return nil, err
	}

	clientID, err := packet.ClientID.ToString()
	if err != nil {
		return nil, err
	}

	userName, err := packet.UserName.ToString()
	if err != nil {
		return nil, err
	}
	return models.ResponseOk{
		ReturnCode:  packet.ReturnCode,
		ErrorString: errString,
		ClientID:    clientID,
		ClientType:  packet.ClientType,
		UserName:    userName,
		ExpiresIn:   packet.ExpiresIn,
		UserID:      packet.UserID,
	}, nil
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
