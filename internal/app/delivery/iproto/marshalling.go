package iproto

import (
	"bytes"
	"encoding/gob"
)

func Marshal(val interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnMarshal(dst []byte, val interface{}) error {
	buf := bytes.NewBuffer(dst)
	enc := gob.NewDecoder(buf)
	if err := enc.Decode(val); err != nil {
		return err
	}
	return nil
}
