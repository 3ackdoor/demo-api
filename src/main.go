package main

import (
	"github.com/3ackdoor/go-demo-api/src/module/user/controller"
	"github.com/3ackdoor/go-demo-api/src/module/user/service"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router *gin.Engine
}

func NewRoutes(app *gin.Engine) Routes {
	r := Routes{
			router: app,
	}

	v1 := r.router.Group("/v1")

	controller.NewUserController(v1, service.NewUserService())

	return r
}

func (r Routes) Run(addr ...string) error {
	return r.router.Run()
}

func main() {
	app := gin.New();
	app.Use(gin.Logger(), gin.Recovery())

	r := NewRoutes(app)
	r.Run()

}