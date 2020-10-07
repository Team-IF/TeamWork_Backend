package db

import (
	"time"

	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	"github.com/Team-IF/TeamWork_Backend/utils"
)

func CreateUser(name, displayName, avatar, email, password, verifyCode string) (uint, error) {
	hashedPassword := utils.HashAndSalt(password)

	emailStruct := dbmodels.UserEmail{VerifyCode: verifyCode, Date: time.Now(), Verified: false, Value: email}
	userStruct := dbmodels.User{Name: name, DisplayName: displayName, Avatar: avatar, Email: emailStruct, Password: hashedPassword}

	result := utils.GetDB().Create(&userStruct)

	return userStruct.ID, result.Error
}

func FindUserByID(id uint) (data *dbmodels.User, err error) {
	err = utils.GetDB().Find(&data, id).Error
	return
}

func VerifyEmail(email, verifyCode string) error {
	var data dbmodels.UserEmail
	if err := utils.GetDB().Where("value = ? AND verify_code = ?", email, verifyCode).First(&data).Error; err != nil {
		return err
	}

	if data.Verified {
		return ErrAlreadyVerified
	}

	if time.Now().After(data.Date.Add(time.Hour * 3)) {
		return ErrExpired
	}

	result := utils.GetDB().Model(&dbmodels.UserEmail{}).Where("id = ?", data.ID).Update("verified", true)
	return result.Error
}
