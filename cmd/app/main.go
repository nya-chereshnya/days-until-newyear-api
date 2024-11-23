package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(MW())
	r.GET("/", Handler)
	return r
}

func MW() gin.HandlerFunc {
	return func(c *gin.Context) {
		if val := c.Request.Header.Get("User-Role"); val == "admin" {
			log.Println("red button user detected")
		}
		c.Next()
	}
}

func Handler(c *gin.Context) {
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Until(date)
	c.String(http.StatusOK, "time until 01.01.2025: %d", int64(duration.Hours()/24))
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
