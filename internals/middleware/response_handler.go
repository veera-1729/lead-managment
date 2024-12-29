package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/veera-1729/lead-managment/internals/constants"
)

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Retrieve the response and error set by the handler
		data, exists := c.Get(constants.Response)
		if !exists {
			data = nil
		}
		err, exists := c.Get(constants.Error)
		if !exists {
			err = nil
		}
		// Send JSON response
		c.JSON(200, gin.H{
			"data":   data,
			"status": err == nil,
			"error":  err,
		})

	}
}

func Middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
