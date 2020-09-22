package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// Synchronize database
func MakeMigrations(connectionString string) *gorm.DB {
	var db *gorm.DB
	var err error
	retryCount := 10
	for {
		db, err = gorm.Open("mysql", connectionString)
		if err != nil {
			log.Print("mysql connection error : ", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
		if retryCount == 0 {
			log.Print("maximum retry reached")
			break
		}
		retryCount--
	}
	db.AutoMigrate(&Devices{})
	db.AutoMigrate(&Location{})
	return db
}
