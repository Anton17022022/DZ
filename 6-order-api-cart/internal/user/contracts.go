package user

type Jwt interface{
	Create(phoneNumber, sessionId string) (string, error)
}