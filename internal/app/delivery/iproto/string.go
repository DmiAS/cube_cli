package iproto

import (
	"github.com/DmiAS/cube_cli/internal/app/models"
)

func strToProtoString(str string) models.String {
	length := int32(len(str))
	b := make([]int8, length)
	for i, char := range str {
		b[i] = int8(char)
	}
	return models.String{Len: length, Str: b}
}
