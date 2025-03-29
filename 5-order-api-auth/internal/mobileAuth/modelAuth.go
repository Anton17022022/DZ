package mobileauth

import "gorm.io/gorm"

type MobileAuthUser struct {
	gorm.Model
	Phone     string `validate:"required,e164"`
	SessionId string `validate:"required" gorm:"uniqueIndex"`
	VerifyCode string `validate:"required"`
}

func NewMobileAuthUser(phoneNumber string) *MobileAuthUser{
	user := &MobileAuthUser{
		Phone: phoneNumber,
	}
	user.GenerateSessionId()
	user.GenerateVerifyCode()
	return user
}

func (user *MobileAuthUser) GenerateSessionId() {
	// умная генерация айдишника сессии TODO
	user.SessionId = user.Phone + "f"
}

func (user *MobileAuthUser) GenerateVerifyCode(){
	// умная генерация кода подтверждения TODO
	user.VerifyCode = "1234"
}
