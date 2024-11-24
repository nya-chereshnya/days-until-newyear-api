package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(c *gin.Context) {
	d := e.s.DaysLeft()
	h := gin.H{
		"Status": d,
	}
	c.HTML(http.StatusOK, "index.html", h)
	// c.String(http.StatusOK, "time until 01.01.2025: %d", d)
}
