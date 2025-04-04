package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	Secret string
}

type JwtData struct {
	PhoneNumber string
	SessionId   string
}

func NewJwt(secret string) *Jwt {
	return &Jwt{secret}
}

func (j *Jwt) Create(phoneNumber, sessionId string) (string, error) {
	stringTime := time.Now().Format("2006-01-02 15:04:05")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId":   sessionId,
		"phoneNumber": phoneNumber,
		"time": stringTime,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *Jwt) Parse(token string) (bool, *JwtData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	return true, &JwtData{
		PhoneNumber: t.Claims.(jwt.MapClaims)["phoneNumber"].(string),
		SessionId:   t.Claims.(jwt.MapClaims)["sessionId"].(string),
	}
}
