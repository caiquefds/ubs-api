package models

import (
	"io/ioutil"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&UBS{}, &City{})
	if err != nil {
		return
	}

	c, ioErr := ioutil.ReadFile("/home/caique/Desktop/Repository/ubs-api/initdb_cities.sql")
	if ioErr != nil {
		return
	}

	sql := string(c)

	err = database.Exec(sql).Error

	if err != nil {
		return
	}

	c, ioErr = ioutil.ReadFile("/home/caique/Desktop/Repository/ubs-api/initdb_ubs.sql")
	if ioErr != nil {
		return
	}

	sql = string(c)

	err = database.Exec(sql).Error

	if err != nil {
		return
	}

	DB = database
}
