package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adhyttungga/go-crud-gin-swagger/config"
	_ "github.com/adhyttungga/go-crud-gin-swagger/docs"
	"github.com/adhyttungga/go-crud-gin-swagger/router"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:5000
// @BasePath /api
func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	log.Info().Msg("Server Initialized")
	db := config.ConnectDB()

	routes := router.NewRouter(db)

	server := &http.Server{
		Addr:    ":" + config.Config.ServerPort,
		Handler: routes,
	}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().AnErr("Listen: ", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().AnErr("Server Shutdown: ", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info().Msg("Timeout of 5 seconds.")
	}
	log.Info().Msg("Server Exiting")
}
