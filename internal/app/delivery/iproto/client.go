package iproto

import "github.com/DmiAS/cube_cli/internal/app/models"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func Send(token, scope string) {
	_, _ = packRequest(token, scope)
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

	return Marshal(models.RequestBody{
		SvcMsg: svcMsg,
		Token:  svcToken,
		Scope:  svcScope,
	})
}
