package models

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	SvcID      int32
	BodyLength int32
	RequestID  int32
}

func (h *Header) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, h); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (h *Header) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.LittleEndian, h); err != nil {
		return err
	}
	return nil
}
