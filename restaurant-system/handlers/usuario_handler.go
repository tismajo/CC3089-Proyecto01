package handlers

import (
	"context"
	"net/http"
	"time"

	"restaurant-system/config"
	"restaurant-system/models"
	"restaurant-system/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuario.FechaRegistro = time.Now()

	_, err := config.DB.Collection("usuarios").InsertOne(context.Background(), usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

func GetUsuarios(c *gin.Context) {

	cursor, err := config.DB.Collection("usuarios").Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var usuarios []models.Usuario
	cursor.All(context.Background(), &usuarios)

	c.JSON(http.StatusOK, usuarios)
}

func BulkUsuarios(c *gin.Context) {

	err := services.BulkInsertUsuarios()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Bulk insert ejecutado"})
}