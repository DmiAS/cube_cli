package models

import (
	"bytes"
	"encoding/binary"
)

type Response struct {
	returnCode  int32
	errorString String
	clientID    String
	clientType  int32
	userName    String
	expiresIn   int32
	userID      int64
}

type ResponsePacket struct {
	Header Header
	Body   Response
}

func (r *Response) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, r.returnCode); err != nil {
		return nil, err
	}

	errorString, err := r.errorString.Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(errorString); err != nil {
		return nil, err
	}

	clientID, err := r.clientID.Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(clientID); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, r.clientType); err != nil {
		return nil, err
	}

	userName, err := r.userName.Marshal()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(userName); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, r.expiresIn); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, r.userID); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (r *Response) UnMarshal(data []byte) error {
	buf := bytes.NewBuffer(data)
	if err := binary.Read(buf, binary.LittleEndian, &r.returnCode); err != nil {
		return err
	}

	if err := r.errorString.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	_ = buf.Next(r.errorString.Length())

	if err := r.clientID.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	_ = buf.Next(r.clientID.Length())

	if err := binary.Read(buf, binary.LittleEndian, &r.clientType); err != nil {
		return err
	}

	if err := r.userName.UnMarshal(buf.Bytes()); err != nil {
		return err
	}
	_ = buf.Next(r.userName.Length())

	if err := binary.Read(buf, binary.LittleEndian, &r.expiresIn); err != nil {
		return err
	}

	if err := binary.Read(buf, binary.LittleEndian, &r.userID); err != nil {
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
