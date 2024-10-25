package main

import (
	"github.com/gin-gonic/gin"
)



func (app *Config) RegisterRoutes() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        // Add routes here
        api.GET("/health", app.HealthCheck)
        api.GET("/example", app.Example)
    }

    return r
}
