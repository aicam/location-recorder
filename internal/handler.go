package internal

import (
	"github.com/aicam/location-recorder/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var bodyJson struct {
			DeviceID string `json:"device_id"`
			Password string `json:"password"`
		}
		err := context.BindJSON(&bodyJson)
		if err != nil {
			WrongParameterRes(context)
			return
		}

	}
}

func (s *Server) RecordLoc() gin.HandlerFunc {
	return func(context *gin.Context) {
		type DeviceInfo struct {
			Longitude float32 `json:"longitude"`
			Latitude  float32 `json:"latitude"`
			Altitude  float32 `json:"altitude"`
			MacAddr   string  `json:"mac_addr"`
			RSSI      int     `json:"rssi"`
		}
		var bodyJson struct {
			NumberOfDevices uint         `json:"number_of_devices"`
			DeviceID        string       `json:"device_id"`
			Devices         []DeviceInfo `json:"devices"`
		}
		err := context.BindJSON(&bodyJson)
		if err != nil {
			WrongParameterRes(context)
		}
		for _, dev := range bodyJson.Devices {
			loc := database.Location{
				DeviceID:  bodyJson.DeviceID,
				Longitude: dev.Longitude,
				Latitude:  dev.Latitude,
				Altitude:  dev.Altitude,
				MacAddr:   dev.MacAddr,
				RSSI:      dev.RSSI,
			}
			s.DB.Save(&loc)
		}
		context.JSON(http.StatusOK, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "Ok", Message: "Location recorded successfully"})
	}
}
