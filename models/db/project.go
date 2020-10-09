package db

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255)"`
	Description  string
	Password     string `gorm:"type:varchar(255)"`
	ProjectOwner uint
	Owned        bool `gorm:"-"`
}

type ProjectMemeber struct {
	ProjectID uint
	UserID    uint
}
