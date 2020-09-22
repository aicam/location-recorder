package main

import (
	"github.com/aicam/location-recorder/internal"
	"github.com/aicam/location-recorder/internal/database"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

const DatabaseConnectionString = "aicam:021021ali@tcp(127.0.0.1:3306)/beacon?charset=utf8mb4&loc=Local&parseTime=True"

func main() {
	err := os.Setenv("IR", "Iran/Tehran")
	if err != nil {
		log.Print("Problem in setting timezone")
	}
	// initialize new server with db and router
	s := internal.NewServer()
	// initialize database
	db := database.MakeMigrations(DatabaseConnectionString)
	s.DB = db
	s.DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	s.Route()
	err = http.ListenAndServe("0.0.0.0:2020", s.Router)
	if err != nil {
		log.Print(err)
	}

}
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				_ = createdAtField.Set(now)
			}
		}
		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				_ = updatedAtField.Set(now)
			}
		}
	}
}
