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
