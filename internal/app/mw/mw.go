package mw

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const RoleAdmin = "admin"

func RoleCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if val := c.Request.Header.Get("User-Role"); strings.EqualFold(RoleAdmin, val) {
			log.Println("red button user detected")
		}
		c.Next()
	}
}
