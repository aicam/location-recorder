package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrongParameterRes(c *gin.Context) {
	c.JSON(http.StatusBadRequest, struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{Code: 304, Message: "Wrong parameter!!"})
}
