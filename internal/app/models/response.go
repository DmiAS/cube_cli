package models

type OkResponseBody struct {
	returnCode int32
	clientID   String
	clientType int32
	userName   String
	expiresIn  int32
	userID     int64
}

type ErrResponseBody struct {
	returnCode  int32
	errorString String
}

type Response struct {
	header  Header
	okResp  OkResponseBody
	errResp ErrResponseBody
}
