package internal

import "github.com/gin-gonic/gin"

func WrongParameterRes(c *gin.Context) {
	c.JSON()
}
