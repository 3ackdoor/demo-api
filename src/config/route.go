package config

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/controller"
	"github.com/3ackdoor/go-demo-api/src/module/user/service"
	"github.com/gin-gonic/gin"
)

type Route struct {
	router *gin.Engine
}

func NewRoutes(app *gin.Engine) Route {
	r := Route{
			router: app,
	}

	v1 := r.router.Group("/v1")

	controller.NewUserController(v1, service.NewUserService())

	return r
}

func (r Route) Run(addr ...string) error {
	return r.router.Run()
}
