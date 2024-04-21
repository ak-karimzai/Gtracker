// Package main is the entry point for the Gtracker API server.
// This server provides endpoints for tracking and managing goals.
// The server handles user authentication, database operations, and request routing.
package main

import (
	"fmt"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/docs"
	_ "git.iu7.bmstu.ru/ka19iu10/Gtracker/docs"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.iu7.bmstu.ru/ka19iu10/Gtracker/cmd/server"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/handler"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/repository"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/internal/service"
	auth_token "git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/auth-token"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/db"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/logger"
	"git.iu7.bmstu.ru/ka19iu10/Gtracker/pkg/util"
	"golang.org/x/net/context"
)

var config util.Config

// @title Goal tracker
// @version 0.1
// @description API Server for Goal tracker app

// @host localhost
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	lgr, err := logger.NewLogger(config.LoggerFilePath)
	if err != nil {
		log.Fatal(err)
	}

	tokenMaker, err := auth_token.NewJWTToken(config.TokenSecretKey, config.TokenValidationTime)
	if err != nil {
		lgr.Fatal(err)
	}

	lgr.Info("migrating database")
	if err = db.Migrate(config.MigrationUrl,
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBName,
		config.SSLMode,
		config.DBPassword); err != nil {
		lgr.Fatal(err)
	}

	conn, err := db.NewPSQL(config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBName,
		config.SSLMode,
		config.DBPassword,
	)
	if err != nil {
		lgr.Fatalf("error while connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(conn, lgr)
	services := service.NewService(repos, tokenMaker, lgr)
	handlers := handler.NewHandler(services, tokenMaker, lgr)

	srv := new(server.Server)
	go func() {
		time.Sleep(time.Second)
		if err := srv.Run(config.ServerPort, handlers.InitRoutes(config.BasePath)); err != nil {
			lgr.Fatalf("an error occured during start of server: %s", err)
		}
	}()
	lgr.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	lgr.Info("http server closes connections")

	if err := srv.Close(context.Background()); err != nil {
		lgr.Errorf("an error occured during closing connection with http server: %s", err.Error())
	}

	conn.Close()
}

// init initializes the application configuration and sets up the Swagger info.
//
// It loads the configuration using the util.NewConfig() function. If an error occurs during loading,
// it logs the error and exits the application.
// The loaded configuration is then printed to the log.
// If the ServerPort in the configuration is not set to "80", the SwaggerInfo.Host is updated to include the ServerPort.
// If the BasePath in the configuration is not empty, the SwaggerInfo.BasePath is updated.
func init() {
	var err error
	config, err = util.NewConfig()
	if err != nil {
		log.Fatalf("error while loading configs: %s", err.Error())
	}

	log.Print(config)
	if config.ServerPort != "80" {
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", config.ServerPort)
	}
	if config.BasePath != "" {
		docs.SwaggerInfo.BasePath = config.BasePath
	}
}
