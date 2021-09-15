package models

type Request struct {
	SvcMsg int32
	Token  String
	Scope  String
}

type RequestPacket struct {
	header Header
	body   Request
}
