package models

import "testing"

func stringToString(str string) String {
	length := len(str)
	b := make([]int8, length)
	for i := range str {
		b[i] = int8(str[i])
	}
	return String{
		Len: int32(length),
		Str: b,
	}
}

func compareSlices(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func compareStrings(a, b String) bool {
	return a.Len == b.Len && compareSlices(a.Str, b.Str)
}

func compareReq(a, b Request) bool {
	return a.SvcMsg == b.SvcMsg && compareStrings(a.Token, b.Token) && compareStrings(a.Scope, b.Scope)
}

func compareHeaders(a, b Header) bool {
	return a.BodyLength == b.BodyLength && a.RequestID == b.RequestID && a.SvcID == b.SvcID
}

func TestStringCoding(t *testing.T) {
	s := stringToString("token")

	data, err := s.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newS String

	if err := newS.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareStrings(s, newS) {
		t.Fatalf("%v != %v", newS, s)
	}
}

func TestStringZero(t *testing.T) {
	s := stringToString("")

	data, err := s.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newS String

	if err := newS.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareStrings(s, newS) {
		t.Fatalf("%v != %v", newS, s)
	}
}

func TestStringOne(t *testing.T) {
	s := stringToString("t")

	data, err := s.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newS String

	if err := newS.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareStrings(s, newS) {
		t.Fatalf("%v != %v", newS, s)
	}
}

func TestRequestCoding(t *testing.T) {
	token := stringToString("token")
	scope := stringToString("scope")
	r := Request{
		SvcMsg: 10,
		Token:  token,
		Scope:  scope,
	}

	data, err := r.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newR Request
	if err := newR.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareReq(newR, r) {
		t.Fatalf("%v != %v", newR, r)
	}
}

func TestRequestZero(t *testing.T) {
	r := Request{
		SvcMsg: 10,
	}

	data, err := r.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newR Request
	if err := newR.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareReq(newR, r) {
		t.Fatalf("%v != %v", newR, r)
	}
}

func TestHeaderCoding(t *testing.T) {
	h := Header{
		SvcID:      10,
		BodyLength: 20,
		RequestID:  30,
	}

	data, err := h.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newH Header

	if err := newH.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareHeaders(h, newH) {
		t.Fatalf("%v != %v", newH, h)
	}
}
