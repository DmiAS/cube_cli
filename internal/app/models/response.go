package models

import (
	"bytes"
	"encoding/binary"
)

type ResponseOk struct {
	ReturnCode  int32
	ErrorString string
	ClientID    string
	ClientType  int32
	UserName    string
	ExpiresIn   int32
	UserID      int64
}

type ResponseErr struct {
	ReturnCode  int32
	ErrorString string
}

type ResponseBody struct {
	ReturnCode  int32
	ErrorString String
	ClientID    String
	ClientType  int32
	UserName    String
	ExpiresIn   int32
	UserID      int64
}

type ResponsePacket struct {
	Header Header
	Body   ResponseBody
}

func (r *ResponseBody) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, r.ReturnCode); err != nil {
		return nil, err
	}

	ErrorString, err := r.ErrorString.Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(ErrorString); err != nil {
		return nil, err
	}

	ClientID, err := r.ClientID.Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(ClientID); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, r.ClientType); err != nil {
		return nil, err
	}

	UserName, err := r.UserName.Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(UserName); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, r.ExpiresIn); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, r.UserID); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (r *ResponseBody) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.LittleEndian, &r.ReturnCode); err != nil {
		return err
	}

	if err := r.ErrorString.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	_ = buf.Next(r.ErrorString.Length())

	if err := r.ClientID.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	_ = buf.Next(r.ClientID.Length())

	if err := binary.Read(buf, binary.LittleEndian, &r.ClientType); err != nil {
		return err
	}

	if err := r.UserName.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	_ = buf.Next(r.UserName.Length())

	if err := binary.Read(buf, binary.LittleEndian, &r.ExpiresIn); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &r.UserID); err != nil {
		return err
	}

	return nil
}

func (r *ResponsePacket) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, r.Header); err != nil {
		return nil, err
	}

	data, err := r.Body.Marshal()
	if err != nil {
		return nil, err
	}

	return append(buf.Bytes(), data...), nil
}

func (r *ResponsePacket) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)

	if err := binary.Read(buf, binary.LittleEndian, &r.Header); err != nil {
		return err
	}

	if err := r.Body.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
