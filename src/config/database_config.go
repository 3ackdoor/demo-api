package config

import (
	"fmt"
	"log"

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
	timezone = "UTC"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s TimeZone=%s",
		host,
		port,
		user,
		password,
		dbname,
		schema,
		timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true, PrepareStmt: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Error occurred when try to connect to database:", err)
	}
	return db
}