package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(255);uniqueIndex"`
	Email       UserEmail `gorm:"foreignkey:UserID"`
	DisplayName string    `gorm:"varchar(255)"`
	Avatar      string    `gorm:"varchar(255)"`
	Password    string    `gorm:"varchar(255)"`
	Status      int       `gorm:"type:tinyint(10);default:1"`
}

type UserEmail struct {
	gorm.Model
	UserID     uint
	VerifyCode string `gorm:"type:varchar(100);uniqueIndex"`
	Date       time.Time
	Verified   bool
	Value      string `gorm:"type:varchar(255);unique_index"`
}
