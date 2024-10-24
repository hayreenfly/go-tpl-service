package main

import (
	"database/sql"
	"fmt"
	"go-tpl-service/data"
	_ "go-tpl-service/docs"
	"go-tpl-service/internal/config"
	"log"
	"time"

	"github.com/gin-gonic/gin"
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

var DB *sql.DB
var counts int64

// @title Go Template Service API
// @version 1.0
// @description This is a Go Template Service
// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to DB
	conn := connectToDB(cfg.Database.GetDSN())
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// Set up database config
	app := Config{
		DB:     conn,
		Models: data.New(conn), // Assuming data.New() initializes your models
	}

	r := gin.Default()

	// register routes
	app.routes(r)

	// Swagger route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting server on port %s", port)

	// Run the server
	r.Run(port) // Listen and serve
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// connectToDB initializes the database connection
func connectToDB(dsn string) *sql.DB {
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}