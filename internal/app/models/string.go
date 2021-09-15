package models

import (
	"bytes"
	"encoding/binary"
)

const int32Size = 4

type String struct {
	Len int32
	Str []int8
}

func (s String) Length() int {
	return int32Size + len(s.Str)
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

	for i := int32(0); i < s.Len; i++ {
		char, err := buf.ReadByte()
		if err != nil {
			return err
		}
		s.Str = append(s.Str, int8(char))
	}

	return nil
}
