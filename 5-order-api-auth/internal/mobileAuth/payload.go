package mobileauth

type SessionIdRequest struct {
}

type SessionIdResponse struct {
	SessionId string `json:"sessionid"`
}

type MobileVerifyRequest struct {
	SessionId string `json:"sessionid"`
	Code      string `json:"code"`
}

type MobileVerifyResponse struct{
	Token string `json:"token"`
}

type SomeUsefullRequest struct{
}

type SomeUsefullResponse struct {
	PhoneNumber string `json:"phone_number"`
}
