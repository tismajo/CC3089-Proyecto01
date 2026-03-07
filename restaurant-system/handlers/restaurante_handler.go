package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"restaurant-system/services"
	"restaurant-system/config"
	"restaurant-system/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRestaurante(c *gin.Context) {
	var restaurante models.Restaurante

	if err := c.ShouldBindJSON(&restaurante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurante.FechaRegistro = time.Now()

	_, err := config.DB.Collection("restaurantes").InsertOne(context.Background(), restaurante)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, restaurante)
}

func GetRestaurantes(c *gin.Context) {

	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	skip, _ := strconv.ParseInt(c.DefaultQuery("skip", "0"), 10, 64)

	opts := options.Find().
		SetLimit(limit).
		SetSkip(skip).
		SetSort(bson.D{{"fechaRegistro", -1}}).
		SetProjection(bson.M{
			"nombre":      1,
			"descripcion": 1,
			"categorias":  1,
		})

	cursor, err := config.DB.Collection("restaurantes").
		Find(context.Background(), bson.M{}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var restaurantes []models.Restaurante
	cursor.All(context.Background(), &restaurantes)

	c.JSON(http.StatusOK, restaurantes)
}

func RestaurantesCercanos(c *gin.Context) {

	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lng, _ := strconv.ParseFloat(c.Query("lng"), 64)
	dist, _ := strconv.Atoi(c.DefaultQuery("dist", "5000"))

	resultados, err := services.RestaurantesCercanosService(lat, lng, dist)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resultados)
}