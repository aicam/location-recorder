package database

import "github.com/jinzhu/gorm"

type Devices struct {
	gorm.Model
	Password string `json:"password"`
	DeviceID string `json:"device_id"`
}

type Location struct {
	gorm.Model
	DeviceID uint `json:"device_id"`
	Longitude float32 `json:"longitude"`
	Latitude float32 `json:"latitude"`
	Altitude float32 `json:"altitude"`
}