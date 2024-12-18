package main

import (
	"database/sql"
	"fmt"
	"go-tpl-service/data"
	database "go-tpl-service/db"
	_ "go-tpl-service/docs"
	"go-tpl-service/internal/config"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)


type Config struct {
	DB     *sql.DB
	Models data.Models
}


// @title Go Template Service API
// @version 1.0
// @description This is a Go Template Service
// @host localhost:8080
// @BasePath /api
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to DB
	conn := database.ConnectDB(cfg.Database.GetDSN())
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// Set up database config
	app := Config{
		DB:     conn,
		Models: data.New(conn), // Assuming data.New() initializes your models
	}

	// register routes
	r := app.RegisterRoutes()

	// Swagger route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting server on port %s", port)

	// Run the server
	r.Run(port) // Listen and serve
}

