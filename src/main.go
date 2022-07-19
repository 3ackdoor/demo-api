package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/3ackdoor/go-demo-api/src/config"
	"github.com/3ackdoor/go-demo-api/src/middleware"
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
	timezone = "UTC"
)

func main() {

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
		log.Fatal("error occurred when trying to connect to database:", err)
	}

	app := gin.New()
	app.Use(
		gin.Logger(),
		middleware.ResponseHandler(),
		gin.CustomRecovery(middleware.ErrorHandler()),
	)

	r := config.NewRoutes(app, db)
	srv := run(r)
	gracefulShutdown(srv)

}

func run(r config.Route) *http.Server {
	var port string
	if p := os.Getenv("port"); p != "" {
		port = p
	} else {
		port = "8080" // default port
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println()
	log.Printf("Listening and serving HTTP on :%s\n", port)
	return srv
}

// ref: https://github.com/hlandau/service/issues/10 and https://github.com/golang/go/issues/9463
func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	killSignal := <-quit
	switch killSignal {
	case syscall.SIGINT:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server gracefully stopped")
}
