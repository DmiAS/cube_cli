package iproto

import (
	"errors"
	"fmt"
	"testing"

	cli2 "github.com/DmiAS/cube_cli/internal/app/cli"
	"github.com/DmiAS/cube_cli/internal/app/connection"
	"github.com/DmiAS/cube_cli/internal/app/mocks"
	"github.com/DmiAS/cube_cli/internal/app/models"
)

func closeFunc(err error) mocks.ErrFn {
	return func() error {
		return err
	}
}

func writeFunc(withError bool) mocks.WriteFn {
	return func([]byte) (int, error) {
		var err error = nil
		if withError {
			err = errors.New("new error")
		}
		return 0, err
	}
}

func dialFunc(conn connection.Connection, err error) mocks.DialFn {
	return func() (connection.Connection, error) {
		return conn, err
	}
}

func compareResponseOk(a, b models.ResponseOk) bool {
	return a.UserID == b.UserID && a.ExpiresIn == b.ExpiresIn &&
		a.UserName == b.UserName && a.ClientType == b.ClientType &&
		a.ClientID == b.ClientID && a.ReturnCode == b.ReturnCode
}

func TestClientInvalidDial(t *testing.T) {
	dialFn := dialFunc(nil, errors.New("can't connect to server"))

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	if _, err := cli.Send("token", "scope"); err == nil {
		t.Fatalf("err = %v", err)
	}
}

func TestClientInvalidWrite(t *testing.T) {
	writeFn := writeFunc(true)

	closeFn := closeFunc(nil)

	conn := &mocks.Connection{
		CloseWriteFn: closeFn,
		WriteFn:      writeFn,
	}

	dialFn := dialFunc(conn, nil)

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	if _, err := cli.Send("token", "scope"); err == nil {
		t.Fatalf("err = %v", err)
	}
}

func TestClientInvalidCloseWrite(t *testing.T) {
	writeFn := writeFunc(false)
	closeFn := closeFunc(errors.New("can't close connection"))

	conn := &mocks.Connection{
		CloseWriteFn: closeFn,
		WriteFn:      writeFn,
	}

	dialFn := dialFunc(conn, nil)

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	if _, err := cli.Send("token", "scope"); err == nil {
		t.Fatalf("err = %v", err)
	}
}

func TestClientReadErr(t *testing.T) {
	writeFn := writeFunc(false)
	closeFn := closeFunc(nil)

	errString := "bad scope"
	readFn := func(_ []byte) ([]byte, error) {
		resp := models.ResponsePacket{
			Header: models.Header{},
			Body: models.ResponseBody{
				ReturnCode:  1,
				ErrorString: strToProtoString(errString),
			},
		}

		return resp.Marshal()
	}

	conn := &mocks.Connection{
		CloseWriteFn: closeFn,
		ReadFn:       readFn,
		WriteFn:      writeFn,
	}

	dialFn := dialFunc(conn, nil)

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	resp, err := cli.Send("token", "scope")
	if err != nil {
		t.Fatalf("err = %v", err)
	}

	respErr, ok := resp.(models.ResponseErr)
	if !ok {
		t.Fatal("invalid type of struct, expected ResponseErr")
	}

	if respErr.ErrorString != errString {
		t.Fatalf("err string = %s", respErr.ErrorString)
	}
}

func TestClientReadOk(t *testing.T) {
	ans := models.ResponseOk{
		ReturnCode: 0,
		ClientID:   "test_client_id",
		ClientType: 2002,
		UserName:   "testuser@mail.ru",
		ExpiresIn:  3600,
		UserID:     101010,
	}

	readFn := func(_ []byte) ([]byte, error) {
		resp := models.ResponsePacket{
			Header: models.Header{},
			Body: models.ResponseBody{
				ReturnCode: ans.ReturnCode,
				ClientID:   strToProtoString(ans.ClientID),
				ClientType: ans.ClientType,
				UserName:   strToProtoString(ans.UserName),
				ExpiresIn:  ans.ExpiresIn,
				UserID:     ans.UserID,
			},
		}

		return resp.Marshal()
	}

	conn := &mocks.Connection{
		CloseWriteFn: closeFunc(nil),
		ReadFn:       readFn,
		WriteFn:      writeFunc(false),
	}

	dialFn := dialFunc(conn, nil)

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	resp, err := cli.Send("token", "scope")
	if err != nil {
		t.Fatalf("err = %v", err)
	}

	respOk, ok := resp.(models.ResponseOk)
	if !ok {
		t.Fatal("invalid type of struct, expected ResponseOk")
	}

	if !compareResponseOk(respOk, ans) {
		t.Fatalf("%v != %v", respOk, ans)
	}

}

func TestClientWriteFields(t *testing.T) {
	token, scope := "abracadabra", "test"

	// здесь проверяю правильно упаковались данные или нет
	readFn := func(data []byte) ([]byte, error) {
		req := &models.RequestPacket{}
		if err := req.UnMarshal(data); err != nil {
			return nil, err
		}

		tokenR, err := req.Body.Token.ToString()
		if err != nil {
			return nil, err
		}
		scopeR, err := req.Body.Scope.ToString()
		if err != nil {
			return nil, err
		}

		if tokenR != token {
			return nil, fmt.Errorf("request token(%s) != token(%s)", tokenR, token)
		}

		if scope != scopeR {
			return nil, fmt.Errorf("request token(%s) != token(%s)", tokenR, token)
		}

		return data, nil
	}
	conn := &mocks.Connection{
		CloseWriteFn: closeFunc(errors.New("can't close connection")),
		ReadFn:       readFn,
	}

	dialFn := dialFunc(conn, nil)

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	if _, err := cli.Send(token, scope); err == nil {
		t.Fatalf("err = %v", err)
	}
}

func TestClientRequestHeader(t *testing.T) {
	token, scope := "abracadabra", "test"

	// здесь проверяю правильно упаковались данные или нет
	readFn := func(data []byte) ([]byte, error) {
		req := &models.RequestPacket{}
		length := req.Header.BodyLength

		body, _ := req.Body.Marshal()
		bodyLength := int32(len(body))
		if length != bodyLength {
			return nil, fmt.Errorf("header length(%d) != body length(%d)", length, bodyLength)
		}

		return data, nil
	}
	conn := &mocks.Connection{
		CloseWriteFn: closeFunc(errors.New("can't close connection")),
		ReadFn:       readFn,
	}

	dialFn := dialFunc(conn, nil)

	connector := mocks.NewConnector(dialFn)
	proto := NewClient(connector)
	cli := cli2.NewCubeClient(proto)

	if _, err := cli.Send(token, scope); err == nil {
		t.Fatalf("err = %v", err)
	}
}
