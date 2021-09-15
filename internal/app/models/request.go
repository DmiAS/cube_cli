package models

import (
	"bytes"
	"encoding/binary"
)

type Request struct {
	SvcMsg int32
	Token  String
	Scope  String
}

type RequestPacket struct {
	Header Header
	Body   Request
}

func (r Request) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, r.SvcMsg); err != nil {
		return nil, err
	}

	token, err := r.Token.Marshal()
	if err != nil {
		return nil, err
	}

	scope, err := r.Token.Marshal()
	if err != nil {
		return nil, err
	}

	if _, err := buf.Write(token); err != nil {
		return nil, err
	}

	if _, err := buf.Write(scope); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
