package models

import "testing"

func stringToSvcString(str string) []int8 {
	b := make([]int8, len(str))
	for i := range str {
		b[i] = int8(str[i])
	}
	return b
}

func slisesIsEqual(a, b []int8) bool {
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

func TestStringCoding(t *testing.T) {
	str := stringToSvcString("token")
	s := &String{
		Len: int32(len(str)),
		Str: str,
	}

	data, err := s.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	newS := new(String)

	if err := newS.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if newS.Len != newS.Len || !slisesIsEqual(s.Str, newS.Str) {
		t.Fatalf("%v != %v", newS, s)
	}
}

func TestStringZero(t *testing.T) {
	str := stringToSvcString("")
	s := &String{
		Len: int32(len(str)),
		Str: str,
	}

	data, err := s.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	newS := new(String)

	if err := newS.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if newS.Len != newS.Len || !slisesIsEqual(s.Str, newS.Str) {
		t.Fatalf("%v != %v", newS, s)
	}
}

func TestStringOne(t *testing.T) {
	str := stringToSvcString("t")
	s := &String{
		Len: int32(len(str)),
		Str: str,
	}

	data, err := s.Marshal()
	if err != nil {
		t.Fatal("error in marshal", err)
	}

	newS := new(String)

	if err := newS.UnMarshal(data); err != nil {
		t.Fatal("error in unmarshal", err)
	}

	if newS.Len != newS.Len || !slisesIsEqual(s.Str, newS.Str) {
		t.Fatalf("%v != %v", newS, s)
	}
}
