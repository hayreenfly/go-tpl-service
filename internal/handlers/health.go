package handlers

import (
	"go-tpl-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the health status of the service
// @Tags Health
// @Success 200 {object} models.JsonResponse
// @Router /health [get]
func HealthCheck(c *gin.Context) {
    response := models.JsonResponse{
        Error:   false,
        Message: "Service is healthy",
    }
    c.JSON(http.StatusOK, response)
}