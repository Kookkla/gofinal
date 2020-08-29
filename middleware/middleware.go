package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	fmt.Println("start #middleware")

	token := c.GetHeader("Authorization")
	if token != "November 10, 2009" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Status code is 401 Unauthorized"})
		c.Abort()
		return
	}

	c.Next()

	fmt.Println("end #middleware")
}
