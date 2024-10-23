package routes

import (
	"go-tpl-service/database"
	"go-tpl-service/internal/handlers"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(router *gin.Engine, dbConfig *database.DBConfig) {
    // Add routes here
    router.GET("/health", handlers.HealthCheck)
    router.GET("/example", handlers.Example)
}
