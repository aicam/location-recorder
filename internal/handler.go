package internal

import "github.com/gin-gonic/gin"



func (s *Server) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var bodyJson struct{
			DeviceID string `json:"device_id"`
			Password string `json:"password"`
		}
		err := context.BindJSON(&bodyJson)

	}
}
