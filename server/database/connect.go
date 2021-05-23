package database

import (
	"crypto.hendrylim/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("user1:password@/covid_server?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to database")
	}

	DB = database

	database.AutoMigrate(&models.Test{}, &models.User{})
	
}