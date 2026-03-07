package handlers

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"restaurant-system/config"
	"restaurant-system/models"
	"restaurant-system/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func CreateOrden(c *gin.Context) {
	var orden models.Orden

	if err := c.ShouldBindJSON(&orden); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CrearOrdenConTransaccion(orden)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Orden creada con transacción"})
}

func GetOrdenes(c *gin.Context) {

	cursor, err := config.DB.Collection("ordenes").
		Find(context.Background(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var ordenes []models.Orden
	cursor.All(context.Background(), &ordenes)

	c.JSON(http.StatusOK, ordenes)
}

func UpdateOrdenEstado(c *gin.Context) {
	id := c.Param("id")

	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := config.DB.Collection("ordenes").UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"estado": "cancelado"}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden actualizada"})
}

func UpdateManyOrdenes(c *gin.Context) {

	_, err := config.DB.Collection("ordenes").UpdateMany(
		context.Background(),
		bson.M{"estado": "pendiente"},
		bson.M{"$set": bson.M{"estado": "procesando"}},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Órdenes actualizadas"})
}

func DeleteOrden(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	_, err := config.DB.Collection("ordenes").DeleteOne(
		context.Background(),
		bson.M{"_id": objID},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden eliminada"})
}

func UploadMenu(c *gin.Context) {

	file, _ := c.FormFile("file")
	src, _ := file.Open()
	defer src.Close()

	bucket, _ := gridfs.NewBucket(config.DB)

	uploadStream, _ := bucket.OpenUploadStream(file.Filename)
	defer uploadStream.Close()

	io.Copy(uploadStream, src)

	c.JSON(http.StatusOK, gin.H{"message": "Archivo subido"})
}

func DownloadFile(c *gin.Context) {

	filename := c.Param("filename")

	bucket, _ := gridfs.NewBucket(config.DB)

	var buf bytes.Buffer
	_, err := bucket.DownloadToStreamByName(filename, &buf)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Archivo no encontrado"})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())
}