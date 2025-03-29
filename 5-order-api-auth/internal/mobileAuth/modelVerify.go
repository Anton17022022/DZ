package mobileauth

import "gorm.io/gorm"

type MobileVerifyUser struct {
	gorm.Model
	SessionId  string
	VerifyCode string
	Token      string
}
