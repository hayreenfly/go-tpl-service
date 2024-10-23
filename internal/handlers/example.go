package handlers

import (
	"go-tpl-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Example godoc
// @Summary Example page
// @Description Returns an example message
// @Tags Example
// @Success 200 {object} models.JsonResponse
// @Router /example [get] // Adjusted the router path to match the function name
func Example(c *gin.Context) {
    response := models.JsonResponse{ // Make sure to use the correct casing here
        Error:   false,
        Message: "Welcome to the example route",
        Data:    nil, // You can set data to nil or any other value if needed
    }
    c.JSON(http.StatusOK, response)
}