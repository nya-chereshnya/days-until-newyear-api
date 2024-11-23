package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", Handler)
	return r
}

func Handler(c *gin.Context) {
	//	c.String(http.StatusOK, "test")
	date := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Until(date)
	c.String(http.StatusOK, "time until 01.01.2025: %d", int64(duration.Hours()/24))
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
