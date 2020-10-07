package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name               string `gorm:"type:varchar(255);uniqueIndex"`
	DisplayName        string `gorm:"varchar(255)"`
	Avatar             string `gorm:"varchar(255)"`
	Password           string `gorm:"varchar(255)"`
	PasswordDate       *time.Time
	PasswordVerifyCode string `gorm:"type:varchar(255);unique_index"`
	Status             int    `gorm:"type:tinyint(10);default:1"`
	EmailVerifyCode    string `gorm:"type:varchar(100);uniqueIndex"`
	EmailDate          *time.Time
	EmailVerified      bool
	Email              string `gorm:"type:varchar(255);unique_index"`
}
