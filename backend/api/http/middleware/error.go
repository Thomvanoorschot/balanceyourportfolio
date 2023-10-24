package middleware

import (
	"github.com/gin-gonic/gin"
)

func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err {
			default:
				c.JSON(-1, gin.H{"error": err.Error()})
			}
		}
	}
}
