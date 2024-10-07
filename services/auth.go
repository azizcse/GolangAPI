package services

import (
	"errors"
	internal "ginapi/internal/models"
	"ginapi/internal/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func InitAuthService(db *gorm.DB) *AuthService {
	db.AutoMigrate(&internal.User{})
	return &AuthService{
		db: db,
	}
}

func (a *AuthService) CheckUserIsExist(email *string) bool {
	var user internal.User
	if err := a.db.Where("email=?", email).Find(&user); err != nil {
		return false
	}
	if user.Email == "" {
		return false
	}
	return true
}

func (a *AuthService) Login(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email can't be empty")
	}

	if password == nil {
		return nil, errors.New("password can't be empty")
	}

	var user internal.User

	if err := a.db.Where("email=?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(*password, user.Password) {
		return nil, errors.New("pasword is incorrect")
	}
	return &user, nil
}

func (a *AuthService) Register(email *string, password *string) (*internal.User, error) {
	if email == nil {
		return nil, errors.New("email can't be empty")
	}

	if password == nil {
		return nil, errors.New("password can't be empty")
	}
	if a.CheckUserIsExist(email) {
		return nil, errors.New("user already exist")
	}

	hashPassword, err := utils.HashPassword(*password)
	if err != nil {
		return nil, err
	}

	var user internal.User

	user.Email = *email
	user.Password = hashPassword

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
