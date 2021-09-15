package models

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
	Resp   Response
}
