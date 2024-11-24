package app

import (
	"github.com/days-until-newyear/internal/app/endpoint"
	"github.com/days-until-newyear/internal/app/mw"
	"github.com/days-until-newyear/internal/app/service"
	"github.com/gin-gonic/gin"
	// "runtime"
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
	a.gin.Static("/static", "../../web/static")
	a.gin.LoadHTMLGlob("../../web/templates/*")
	a.gin.Use(mw.RoleCheck(), mw.RequestLogger(), mw.ResponseLogger())
	a.gin.GET("/", a.e.Status)
	//	a.gin.GET("/api/status", a.e.Status) // REST API

	return a
}

func (a *App) Run() {
	a.gin.Run(":8080")
}
