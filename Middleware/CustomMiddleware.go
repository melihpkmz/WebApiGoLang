package Middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func CustomMiddleware(c *gin.Context){
	c.Next()

	err:=c.Errors.Last()
	if err !=nil{
		statusCode := err.Err.Error()
		val, _ := strconv.Atoi(statusCode)
		c.JSON(val,err.Err)
	}





}


