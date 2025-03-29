package mobileauth

import (
	"5-order-api-auth/pkg/db"
	"errors"

	"gorm.io/gorm"
)

type MobileAuthRepository struct {
	*db.Db
}

func NewMobileAuthRepository(Db *db.Db) *MobileAuthRepository {
	return &MobileAuthRepository{Db: Db}
}

func (repo *MobileAuthRepository) WriteSessionId(user *MobileAuthUser) error {
	userAuth, err := repo.FindBySessionId(user.SessionId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result := repo.Db.Create(user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	if userAuth != nil {
		return errors.New("session existed")
	}
	return err
}

func (repo *MobileAuthRepository) WriteToken(user *MobileVerifyUser) (*MobileVerifyUser, error) {
	//TODO проверка наличия сессии и совпадение кода
	userAuth, err := repo.FindBySessionId(user.SessionId)
	if err != nil {
		return nil, err
	}

	if user.VerifyCode != userAuth.VerifyCode {
		return nil, errors.New("invalid verify code")
	}

	result := repo.Db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *MobileAuthRepository) FindBySessionId(sessionId string) (*MobileAuthUser, error) {
	var user MobileAuthUser
	result := repo.Db.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
