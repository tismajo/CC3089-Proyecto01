package handlers

import (
	"context"
	"net/http"

	"restaurant-system/config"
	"restaurant-system/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RestaurantesMejorCalificados(c *gin.Context) {

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":       "$restaurante_id",
				"promedio": bson.M{"$avg": "$calificacion"},
			},
		},
		{
			"$sort": bson.M{"promedio": -1},
		},
		{
			"$lookup": bson.M{
				"from":         "restaurantes",
				"localField":   "_id",
				"foreignField": "_id",
				"as":           "restaurante",
			},
		},
	}

	cursor, err := config.DB.Collection("resenas").
		Aggregate(context.Background(), pipeline)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var resultados []bson.M
	cursor.All(context.Background(), &resultados)

	c.JSON(http.StatusOK, resultados)
}

func VentasPorMes(c *gin.Context) {

	resultados, err := services.VentasPorMesService()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resultados)
}