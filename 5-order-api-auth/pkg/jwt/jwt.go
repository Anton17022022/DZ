package jwt

import "github.com/golang-jwt/jwt/v5"

type Jwt struct {
	Secret string
}

type  JwtData struct {
	PhoneNumber string
	SesionId string
}

func NewJwt(secret string) *Jwt {
	return &Jwt{
		Secret: secret,
	}
}

func (j *Jwt) Create(sessionId, phoneNumber string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionid": sessionId,
		"phonenumber": phoneNumber,
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
	return t.Valid, &JwtData{
		SesionId: t.Claims.(jwt.MapClaims)["sessionid"].(string),
		PhoneNumber: t.Claims.(jwt.MapClaims)["phonenumber"].(string),
	}
}

