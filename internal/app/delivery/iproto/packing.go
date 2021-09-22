package iproto

import "github.com/DmiAS/cube_cli/internal/app/models"

const cube_svc int32 = 0x00000002

// для передачи по сети используем слайс байт, соответственно нужно завернуть токен и скоуп
// в пакет, сформировать структура пакета запроса и преобразовать ее в слайс байт
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

func packHeader(body []byte) ([]byte, error) {
	bodyLength := int32(len(body))
	var reqID int32 = 0
	return Marshal(&models.Header{
		SvcID:      cube_svc,
		BodyLength: bodyLength,
		RequestID:  reqID,
	})
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
