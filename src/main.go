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
)

func main() {
	db := config.NewDatabase()

	app := gin.New()
	app.Use(
		gin.Logger(),
		middleware.ResponseWriterHandler(),
		gin.CustomRecovery(middleware.RecoveryWriterHandler()),
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
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	fmt.Println()
	log.Printf("Listening and serving HTTP on :%s\n", port)

	return srv
}

// See references
//
// https://github.com/hlandau/service/issues/10
//
// https://github.com/golang/go/issues/9463
func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	switch signal := <-quit; signal {
	case syscall.SIGINT:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server gracefully stopped")
}
