package cli

import (
	"fmt"

	"github.com/DmiAS/cube_cli/internal/app/models"
)

const unknown = "UNKNOWN_ERROR"

var codes = [...]string{
	"CUBE_OAUTH2_ERR_OK",
	"CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND",
	"CUBE_OAUTH2_ERR_DB_ERROR",
	"CUBE_OAUTH2_ERR_UNKNOWN_MSG",
	"CUBE_OAUTH2_ERR_BAD_PACKET",
	"CUBE_OAUTH2_ERR_BAD_CLIENT",
	"CUBE_OAUTH2_ERR_BAD_SCOPE",
}

// Protocol предоставляет интерфейс для передачи данных к cube сервису, абстрагируемся от того
// по какому протоколу они передаются, на этом уровне мы лишь отсылаем токен и скоуп, а принимаем
// Response, который может быть либо структурой ответа, либо ошибки
type Response = interface{}
type Protocol interface {
	Send(token, scope string) (interface{}, error)
}

type CubeClient struct {
	proto Protocol
}

func NewCubeClient(proto Protocol) *CubeClient {
	return &CubeClient{proto: proto}
}

func (c *CubeClient) Send(token, scope string) (Response, error) {
	resp, err := c.proto.Send(token, scope)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CubeClient) PrintResponse(resp Response) {
	switch v := resp.(type) {
	case models.ResponseOk:
		printOkResponse(v)
	case models.ResponseErr:
		printErrResponse(v)
	default:
		fmt.Println("unknown response")
	}
}

func printOkResponse(resp models.ResponseOk) {
	fmt.Printf("client_id: %s\n", resp.ClientID)
	fmt.Printf("clint_type: %d\n", resp.ClientType)
	fmt.Printf("expires_in: %d\n", resp.ExpiresIn)
	fmt.Printf("user_id: %d\n", resp.UserID)
	fmt.Printf("username: %s\n", resp.UserName)
}

func printErrResponse(resp models.ResponseErr) {
	fmt.Printf("error: %s\n", codeToString(resp.ReturnCode))
	fmt.Printf("message: %s\n", resp.ErrorString)
}

func codeToString(code int32) string {
	if code >= 0 && code < int32(len(codes)) {
		return codes[int(code)]
	}
	return unknown
}
