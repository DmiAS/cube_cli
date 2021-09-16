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

func (r *Request) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, r.SvcMsg); err != nil {
		return nil, err
	}

	token, err := r.Token.Marshal()
	if err != nil {
		return nil, err
	}

	scope, err := r.Scope.Marshal()
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

func (r *Request) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.LittleEndian, &r.SvcMsg); err != nil {
		return err
	}

	if err := r.Token.UnMarshal(buf.Bytes()); err != nil {
		return err
	}

	// так как я не знаю сколько было считано байт из buf входе UnMarshal
	// я пропускаю эти байты, высчитывая с помощью Length сколько было считано
	_ = buf.Next(r.Token.Length())

	if err := r.Scope.UnMarshal(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func (r *RequestPacket) Marshal() ([]byte, error) {
	header, err := r.Header.Marshal()
	if err != nil {
		return nil, err
	}

	body, err := r.Body.Marshal()
	if err != nil {
		return nil, err
	}

	return append(header, body...), nil
}

func (r *RequestPacket) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.LittleEndian, &r.Header); err != nil {
		return err
	}

	if err := r.Body.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
