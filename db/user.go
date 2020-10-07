package db

import (
	"time"

	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	"github.com/Team-IF/TeamWork_Backend/utils"
)

func CreateUser(name, displayName, avatar, email, password, verifyCode string) (uint, error) {
	hashedPassword := utils.HashAndSalt(password)

	now := time.Now()
	userStruct := dbmodels.User{Name: name, DisplayName: displayName, Avatar: avatar, Email: email, EmailVerified: false, EmailDate: &now, EmailVerifyCode: verifyCode, Password: hashedPassword, PasswordVerifyCode: "", PasswordDate: nil}

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

func UpdateProfile(id uint, name, displayName, avatar string) error {
	result := utils.GetDB().Model(&dbmodels.User{}).Where("id = ?", id).Updates(&dbmodels.User{
		Name:        name,
		DisplayName: displayName,
		Avatar:      avatar,
	})
	return result.Error
}

func UpdateEmail(id uint, email, verifyCode string) error {
	now := time.Now()
	result := utils.GetDB().Model(&dbmodels.User{}).Where("id = ?", id).Updates(&dbmodels.User{
		Email:           email,
		EmailVerifyCode: verifyCode,
		EmailDate:       &now,
	})
	return result.Error
}

func UpdatePassword(id uint, newPassword string) error {
	hashedNew := utils.HashAndSalt(newPassword)
	result := utils.GetDB().Model(&dbmodels.User{}).Where("id = ?", id).Update("password", hashedNew)
	return result.Error
}

func PasswordVerifyCode(email, verifyCode string) error {
	now := time.Now()
	return utils.GetDB().Model(&dbmodels.User{}).Where("email = ?", email).Updates(&dbmodels.User{PasswordVerifyCode: verifyCode, PasswordDate: &now}).Error
}

func UpdatePasswordWithCode(email, password, verifyCode string) error {
	var data dbmodels.User
	if err := utils.GetDB().Select("password_date").Where("email = ? and password_verify_code = ?", email, verifyCode).First(&data).Error; err != nil {
		return err
	}

	if time.Now().After(data.PasswordDate.Add(time.Hour * 3)) {
		return ErrExpired
	}

	hashedNew := utils.HashAndSalt(password)
	now := time.Now()
	result := utils.GetDB().Model(&dbmodels.User{}).Where("id = ?", data.ID).Updates(
		&dbmodels.User{Password: hashedNew, PasswordDate: &now, PasswordVerifyCode: ""})
	return result.Error
}
