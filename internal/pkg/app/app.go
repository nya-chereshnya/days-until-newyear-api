package app

import (
	"fmt"

	"github.com/days-until-newyear/internal/app/endpoint"
	"github.com/days-until-newyear/internal/app/mw"
	"github.com/days-until-newyear/internal/app/service"
	"github.com/days-until-newyear/internal/config"
	"github.com/gin-gonic/gin"
)

// TODO: add test, upgrade logs, add Redis, add Docker

type App struct {
	e       *endpoint.Endpoint
	s       *service.Service
	gin     *gin.Engine
	address string
}

func New() *App {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.gin = gin.Default()
	a.gin.Use(mw.RoleCheck(), mw.RequestLogger(), mw.ResponseLogger())
	a.gin.GET("/", a.e.GetStatus)
	cfg := config.MustLoad()
	a.address = fmt.Sprintf("%s:%s", cfg.IP, cfg.Port)
	return a
}

func (a *App) Run() {
	a.gin.Run(a.address)
}
