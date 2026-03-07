package handlers

import (
	"context"
	"net/http"
	"time"

	"restaurant-system/config"
	"restaurant-system/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateResena(c *gin.Context) {
	var resena models.Resena

	if err := c.ShouldBindJSON(&resena); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resena.Fecha = time.Now()

	_, err := config.DB.Collection("resenas").
		InsertOne(context.Background(), resena)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resena)
}

func GetResenas(c *gin.Context) {

	cursor, err := config.DB.Collection("resenas").
		Find(context.Background(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var resenas []models.Resena
	cursor.All(context.Background(), &resenas)

	c.JSON(http.StatusOK, resenas)
}

func PlatillosMasVendidos(c *gin.Context) {

	pipeline := []bson.M{
		{"$unwind": "$items"},
		{
			"$group": bson.M{
				"_id":      "$items.articulo_id",
				"cantidad": bson.M{"$sum": "$items.cantidad"},
			},
		},
		{"$sort": bson.M{"cantidad": -1}},
	}

	cursor, err := config.DB.Collection("ordenes").
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