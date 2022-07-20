package config

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/controller"
	"github.com/3ackdoor/go-demo-api/src/module/user/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Route struct {
	Router *gin.Engine
}

func NewRoutes(app *gin.Engine, db *gorm.DB) Route {
	r := Route{
		Router: app,
	}

	v1 := r.Router.Group("/v1")

	controller.NewUserController(v1, service.NewUserService(db))

	return r
}

func (r Route) Run(addr ...string) error {
	return r.Router.Run()
}
