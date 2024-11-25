package mw

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const RoleAdmin = "Santa-Claus"

func RoleCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if val := c.Request.Header.Get("User-Role"); strings.EqualFold(RoleAdmin, val) {
			log.Println("Santa Claus detected")
		}
		c.Next()
	}
}
