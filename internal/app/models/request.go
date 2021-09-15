package models

type RequestBody struct {
	SvcMsg int32
	Token  String
	Scope  String
}

type Request struct {
	header Header
	body   RequestBody
}
