package app

import (
	"github.com/days-until-newyear/internal/app/endpoint"
	"github.com/days-until-newyear/internal/app/mw"
	"github.com/days-until-newyear/internal/app/service"
	"github.com/gin-gonic/gin"
)

type App struct {
	e   *endpoint.Endpoint
	s   *service.Service
	gin *gin.Engine
}

func New() *App {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.gin = gin.Default()
	a.gin.Use(mw.RoleCheck())
	a.gin.GET("/", a.e.Status)
	return a
}

func (a *App) Run() {
	a.gin.Run(":8080")
}
