package db

import (
	"time"
)

type User struct {
	ID       int       `gorm:"type:unsigned int;autoIncrement;primaryKey;uniqueIndex;"`
	User_ID  string    `gorm:"type:varchar(255);uniqueIndex;primaryKey"`
	Email    UserEmail `gorm:"embedded;embeddedPrefix:email_"`
	Name     string    `gorm:"varchar(255)"`
	Avatar   string    `gorm:"varchar(255)"`
	Password string    `gorm:"varchar(255)"`
	updateAt time.Time
}

type UserEmail struct {
	ID         int    `gorm:"type:unsigned int;autoIncrement;primaryKey;uniqueIndex;"`
	VerifyCode string `gorm:"type:varchar(100);uniqueIndex"`
	Date       time.Time
	Verified   bool
	Value      string `gorm:"type:varchar(255);unique_index"`
}
