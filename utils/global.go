package utils

import (
	"github.com/Team-IF/TeamWork_Backend/models"
	"gorm.io/gorm"
)

var (
	gConfig *models.Config
	gDB     *gorm.DB
)

func SetConfig(config *models.Config) {
	gConfig = config
}

func GetConfig() *models.Config {
	return gConfig
}

func SetDB(db *gorm.DB) {
	gDB = db
}

func GetDB() *gorm.DB {
	return gDB
}
