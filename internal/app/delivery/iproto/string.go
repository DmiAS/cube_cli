package iproto

import (
	"encoding/base64"

	"github.com/DmiAS/cube_cli/internal/app/models"
)

func strToProtoString(str string) models.String {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	length := int32(len(encoded))
	b := make([]int8, length)
	for i, char := range encoded {
		b[i] = int8(char)
	}
	return models.String{Len: length, Str: b}
}
