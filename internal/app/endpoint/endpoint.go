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

func (e *Endpoint) GetStatus(c *gin.Context) {
	d := e.s.DaysLeft()
	c.IndentedJSON(http.StatusOK, d)
}
