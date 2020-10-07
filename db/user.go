package db

import (
	"time"

	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	"github.com/Team-IF/TeamWork_Backend/utils"
)

func CreateUser(name, displayName, avatar, email, password, verifyCode string) (uint, error) {
	hashedPassword := utils.HashAndSalt(password)

	userStruct := dbmodels.User{Name: name, DisplayName: displayName, Avatar: avatar, Email: email, EmailVerified: false, EmailDate: time.Now(), EmailVerifyCode: verifyCode, Password: hashedPassword}

	result := utils.GetDB().Create(&userStruct)

	return userStruct.ID, result.Error
}

func SignIn(name, password string) (*dbmodels.User, error) {
	var data dbmodels.User
	result := utils.GetDB().Where("name = ? OR email = ?", name, name).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	if !utils.CheckPassword(password, data.Password) {
		return nil, ErrPasswordNotMatch
	}

	return &data, nil
}

func FindUserByID(id uint) (*dbmodels.User, error) {
	var data dbmodels.User
	err := utils.GetDB().First(&data, id).Error
	return &data, err
}

func VerifyEmail(email, verifyCode string) error {
	var data dbmodels.User
	if err := utils.GetDB().Select("verified, email_date").Where("email = ? AND email_verify_code = ?", email, verifyCode).First(&data).Error; err != nil {
		return err
	}

	if data.EmailVerified {
		return ErrAlreadyVerified
	}

	if time.Now().After(data.EmailDate.Add(time.Hour * 3)) {
		return ErrExpired
	}

	result := utils.GetDB().Model(&dbmodels.User{}).Where("id = ?", data.ID).Update("email_verified", true)
	return result.Error
}
