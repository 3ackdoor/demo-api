package main

import (
	"fmt"
	"log"

	"github.com/3ackdoor/go-demo-api/src/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "golang-demo-api"
	schema   = "public"
)

func main() {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, schema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true, PrepareStmt: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Panic(err)
	}

	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())

	r := config.NewRoutes(app, db)
	r.Run()

}
