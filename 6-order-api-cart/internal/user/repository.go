package user

import (
	"6-order-api-cart/pkg/db"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	*db.Db
}

const (
	TokenLifeTime = 20 * time.Minute
)

func NewUseRepository(db *db.Db) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Auth(user *UserAuth, jwt Jwt) (*UserAuth, error) {
	ok, err := repo.CheckUserExisted(user.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if !ok {
		err := repo.CreateUser(user.PhoneNumber)
		if err != nil {
			return nil, err
		}
	}

	userBySessionId, err := repo.FindAuthBySessionId(user.SessionId)
	if err != nil {
		return nil, err
	}

	if userBySessionId != nil {
		if time.Since(userBySessionId.CreatedAt) > TokenLifeTime {
			result := repo.Db.Delete(&UserAuth{}, userBySessionId.ID)

			if result.Error != nil {
				return nil, result.Error
			}

			for userBySessionId.SessionId == user.SessionId && user.CreatedAt.Minute() == time.Now().Minute() {
				user.GenerateSessionId()
			}
		} else {
			return userBySessionId, nil
		}
	}

	user.Token, err = jwt.Create(user.PhoneNumber, user.SessionId)
	if err != nil {
		return nil, err
	}

	result := repo.Db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) FindAuthBySessionId(sessionId string) (*UserAuth, error) {
	var user *UserAuth
	result := repo.Db.First(&user, "session_id = ?", sessionId)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return user, nil
}

func (repo *UserRepository) CreateUser(phoneNumber string) error {
	result := repo.Db.Create(&User{PhoneNumber: phoneNumber})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) CheckUserExisted(phoneNumber string) (bool, error) {
	var user *User
	result := repo.Db.First(&user, "phone_number = ?", phoneNumber)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, nil
}
