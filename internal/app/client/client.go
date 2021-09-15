package client

import (
	"fmt"
	"log"

	"github.com/DmiAS/cube_cli/internal/app/models"
)

var codes = [...]string{
	"CUBE_OAUTH2_ERR_OK",
	"CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND",
	"CUBE_OAUTH2_ERR_DB_ERROR",
	"CUBE_OAUTH2_ERR_UNKNOWN_MSG",
	"CUBE_OAUTH2_ERR_BAD_PACKET",
	"CUBE_OAUTH2_ERR_BAD_CLIENT",
	"CUBE_OAUTH2_ERR_BAD_SCOPE",
}

type Protocol interface {
	Send(token, scope string) (interface{}, error)
}

type CubeClient struct {
	proto Protocol
}

func NewCubeClient(proto Protocol) *CubeClient {
	return &CubeClient{proto: proto}
}

func (c *CubeClient) Send(token, scope string) {
	resp, err := c.proto.Send(token, scope)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp interface{}) {
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
	fmt.Printf("user_id: %d\n", resp.UserID)
	fmt.Printf("username: %s\n", resp.UserName)
}

func printErrResponse(resp models.ResponseErr) {
	fmt.Printf("error: %s\n", codeToString(resp.ReturnCode))
	fmt.Printf("message: %s\n", resp.ErrorString)
}

func codeToString(code int32) string {
	return codes[int(code)]
}
