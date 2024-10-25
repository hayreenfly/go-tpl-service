package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Example godoc
// @Summary Example page
// @Description Returns an example message
// @Tags Example
// @Success 200 {object} jsonResponse
// @Router /example [get] // Adjusted the router path to match the function name
func (app *Config) Example(c *gin.Context) {
    response := jsonResponse{ // Make sure to use the correct casing here
        Error:   false,
        Message: "Welcome to the example route",
        Data:    nil, // You can set data to nil or any other value if needed
    }
    c.JSON(http.StatusOK, response)
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the health status of the service
// @Tags Health
// @Success 200 {object} jsonResponse
// @Router /health [get]
func (app *Config) HealthCheck(c *gin.Context) {
    response := jsonResponse{
        Error:   false,
        Message: "Service is healthy",
    }
    c.JSON(http.StatusOK, response)
}