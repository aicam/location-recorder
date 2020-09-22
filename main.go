package main

import (
	"github.com/aicam/location-recorder/internal"
	"github.com/aicam/location-recorder/internal/database"
)
const DatabaseConnectionString = "aicam:021021ali@tcp(127.0.0.1:3306)/beacon?charset=utf8mb4&parseTime=True"
func main() {
	// initialize new server with db and router
	s := internal.NewServer()
	// initialize database
	db := database.MakeMigrations(DatabaseConnectionString)
	s.DB = db
	s.Routes()
}
