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

func compareReqPackets(a, b RequestPacket) bool {
	return compareHeaders(a.Header, b.Header) && compareReq(a.Body, b.Body)
}

func compareResponses(a, b ResponseBody) bool {
	return a.UserID == b.UserID && a.ExpiresIn == b.ExpiresIn && a.ClientType == b.ClientType &&
		compareStrings(a.UserName, b.UserName) && compareStrings(a.ClientID, b.ClientID) &&
		compareStrings(a.ErrorString, b.ErrorString)
}

func compareResponsePackets(a, b ResponsePacket) bool {
	return compareResponses(a.Body, b.Body) && compareHeaders(a.Header, b.Header)
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

func TestRequestPacketCoding(t *testing.T) {
	token := stringToString("token")
	scope := stringToString("scope")
	r := RequestPacket{
		Header: Header{
			SvcID:      10,
			BodyLength: 20,
			RequestID:  30,
		},
		Body: Request{
			SvcMsg: 1,
			Token:  token,
			Scope:  scope,
		},
	}

	data, err := r.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newR RequestPacket

	if err := newR.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareReqPackets(r, newR) {
		t.Fatalf("%v != %v", newR, r)
	}
}

func TestResponseCoding(t *testing.T) {
	e := stringToString("e")
	c := stringToString("c")
	u := stringToString("u")
	r := ResponseBody{
		ReturnCode:  10,
		ErrorString: e,
		ClientID:    c,
		ClientType:  20,
		UserName:    u,
		ExpiresIn:   30,
		UserID:      40,
	}

	data, err := r.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newR ResponseBody

	if err := newR.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareResponses(r, newR) {
		t.Fatalf("%v != %v", newR, r)
	}
}

func TestResponseZero(t *testing.T) {
	r := ResponseBody{
		ReturnCode: 10,
		ClientType: 20,
		ExpiresIn:  30,
		UserID:     40,
	}

	data, err := r.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newR ResponseBody

	if err := newR.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareResponses(r, newR) {
		t.Fatalf("%v != %v", newR, r)
	}
}

func TestResponsePacketCoding(t *testing.T) {
	body := ResponseBody{
		ReturnCode: 10,
		ClientType: 20,
		ExpiresIn:  30,
		UserID:     40,
	}

	h := Header{
		SvcID:      10,
		BodyLength: 20,
		RequestID:  30,
	}

	r := ResponsePacket{
		Header: h,
		Body:   body,
	}

	data, err := r.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	var newR ResponsePacket

	if err := newR.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if !compareResponsePackets(r, newR) {
		t.Fatalf("%v != %v", newR, r)
	}
}
