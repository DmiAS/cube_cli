package iproto

import "github.com/DmiAS/cube_cli/internal/app/models"

func packHeader(body []byte) ([]byte, error) {
	var svcID int32 = 0
	bodyLength := int32(len(body))
	var reqID int32 = 0
	return Marshal(&models.Header{
		SvcID:      svcID,
		BodyLength: bodyLength,
		RequestID:  reqID,
	})
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
