package handlers

import (
	"net/http"

	"restaurant-system/services"

	"github.com/gin-gonic/gin"
)

func ExplainOrdenes(c *gin.Context) {

	result, err := services.ExplainOrdenesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}