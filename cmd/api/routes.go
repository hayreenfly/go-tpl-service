package main

import (
	"github.com/gin-gonic/gin"
)



func (app *Config) routes(router *gin.Engine) {
    // Add routes here
    router.GET("/health", app.HealthCheck)
    router.GET("/example", app.Example)
}
