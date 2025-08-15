package api

import (
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	// check := false
	// if check {
	// 	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "middleware"})
	// 	return
	// }

	c.Next()
}
