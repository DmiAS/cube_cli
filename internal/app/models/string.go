package models

import (
	"bytes"
	"encoding/binary"
	"strings"
)

const int32Size = 4

type String struct {
	Len int32
	Str []int8
}

// Используется, чтобы знать, сколько байт из буфера было просчитано
func (s String) Length() int {
	return int32Size + len(s.Str)
}

func (s String) ToString() (string, error) {
	builder := new(strings.Builder)
	for _, char := range s.Str {
		if err := builder.WriteByte(byte(char)); err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}

func (s *String) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, s.Len); err != nil {
		return nil, err
	}

	for _, char := range s.Str {
		if err := buf.WriteByte(byte(char)); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (s *String) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.LittleEndian, &s.Len); err != nil {
		return err
	}

	s.Str = make([]int8, 0, s.Len)
	for i := int32(0); i < s.Len; i++ {
		char, err := buf.ReadByte()
		if err != nil {
			return err
		}
		s.Str = append(s.Str, int8(char))
	}

	return nil
}
