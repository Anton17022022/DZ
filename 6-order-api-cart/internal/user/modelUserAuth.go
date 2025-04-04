package user

import (
	"gorm.io/gorm"
)

type UserAuth struct {
	gorm.Model
	PhoneNumber string `json:"phonenumber" gorm:"required,e164"`
	SessionId   string `json:"sessionid" gorm:"required,uniqueIndex"`
	Token       string `json:"token" gorm:"required"`
}

func NewUser(phoneNumber string) *UserAuth {
	user := &UserAuth{
		PhoneNumber: phoneNumber,
	}
	user.GenerateSessionId()
	return user
}

func (user *UserAuth) GenerateSessionId() {
	// Какое-то умное формирование ид сессии
	user.SessionId = user.PhoneNumber + "f"
}
