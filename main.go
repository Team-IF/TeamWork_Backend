package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"time"

	dbmodels "github.com/Team-IF/TeamWork_Backend/models/db"
	v1 "github.com/Team-IF/TeamWork_Backend/routes/v1"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/BurntSushi/toml"
	"github.com/Team-IF/TeamWork_Backend/middlewares"
	"github.com/Team-IF/TeamWork_Backend/models"
	"github.com/Team-IF/TeamWork_Backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// etcInit()
	applyConfig()
	// initDB()
	// startServer()

	utils.SendVefiryMail("qwerqwer", "gangjun2006@gmail.com")
}

func etcInit() {
	rand.Seed(time.Now().Unix())
}

func applyConfig() {
	rawConfig, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatalln("Failed to load config.")
	}

	var config models.Config
	if _, err := toml.Decode(string(rawConfig), &config); err != nil {
		log.Println(err.Error())
		log.Fatalln("Failed to parsing config.")
	}
	utils.SetConfig(&config)
}

func startServer() {
	config := utils.GetConfig().Server
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)

	}
	r := gin.Default()
	r.Use(middlewares.Cors())
	version1 := r.Group("/v1")
	v1.InitRoutes(version1)
	r.Run(":" + strconv.Itoa(config.Port))
}

func initDB() {
	config := utils.GetConfig()
	log.Println("Initializing Database...")
	var db *gorm.DB
	var err error
	if config.Server.UseTestDB {
		dbConfig := config.TestDB
		db, err = gorm.Open(sqlite.Open(dbConfig.Path), &gorm.Config{})

	} else {
		dbConfig := config.DB
		connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Hostname, dbConfig.Port, dbConfig.DBName)
		db, err = gorm.Open(mysql.Open(connectionInfo), &gorm.Config{})
	}

	if err != nil {
		log.Fatalln("Failed to open database.")
	}
	utils.SetDB(db)
	log.Print("Successfully Connected To Database")

	var models = []interface{}{&dbmodels.User{}, &dbmodels.UserEmail{}}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalln("Failed to perform AutoMigrate.")
	}
	log.Print("Successfully performed AutoMigrate")
}
