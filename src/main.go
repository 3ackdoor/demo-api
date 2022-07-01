package main

import (
	"github.com/3ackdoor/go-demo-api/src/config"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New();
	app.Use(gin.Logger(), gin.Recovery())
	

	r := config.NewRoutes(app)
	r.Run()

}
