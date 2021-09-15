package iproto

import "github.com/DmiAS/cube_cli/internal/app/models"

func packHeader(body []byte) ([]byte, error) {
	var svcID int32 = 0
	bodyLength := int32(len(body))
	var reqID int32 = 0
	return Marshal(models.Header{
		SvcID:      svcID,
		BodyLength: bodyLength,
		RequestID:  reqID,
	})
}
