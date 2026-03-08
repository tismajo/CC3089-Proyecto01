package handlers

import (
	"context"
	"net/http"

	"restaurant-system/config"
	"restaurant-system/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// =======================
// CREATE
// =======================
func CreateArticulo(c *gin.Context) {

	var input struct {
		Nombre        string  `json:"nombre"`
		Descripcion   string  `json:"descripcion"`
		Precio        float64 `json:"precio"`
		Categoria     string  `json:"categoria"`
		Disponible    bool    `json:"disponible"`
		RestauranteID string  `json:"restaurante_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restID, err := primitive.ObjectIDFromHex(input.RestauranteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID restaurante inválido"})
		return
	}

	articulo := models.ArticuloMenu{
		Nombre:        input.Nombre,
		Descripcion:   input.Descripcion,
		Precio:        input.Precio,
		Categoria:     input.Categoria,
		Disponible:    input.Disponible,
		RestauranteID: restID,
	}

	_, err = config.DB.Collection("articulosMenu").
		InsertOne(context.Background(), articulo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, articulo)
}


// =======================
// READ (todos)
// =======================
func GetArticulos(c *gin.Context) {

	cursor, err := config.DB.Collection("articulosMenu").
		Find(context.Background(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var articulos []models.ArticuloMenu
	cursor.All(context.Background(), &articulos)

	c.JSON(http.StatusOK, articulos)
}


// =======================
// READ por restaurante
// =======================
func GetArticulosByRestaurante(c *gin.Context) {

	id := c.Param("id")

	restID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	cursor, err := config.DB.Collection("articulosMenu").
		Find(context.Background(), bson.M{"restaurante_id": restID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var articulos []models.ArticuloMenu
	cursor.All(context.Background(), &articulos)

	c.JSON(http.StatusOK, articulos)
}


// =======================
// UPDATE ONE
// =======================
func UpdateArticulo(c *gin.Context) {

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var body bson.M
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = config.DB.Collection("articulosMenu").
		UpdateOne(context.Background(),
			bson.M{"_id": objID},
			bson.M{"$set": body},
		)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artículo actualizado"})
}


// =======================
// DELETE ONE
// =======================
func DeleteArticulo(c *gin.Context) {

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = config.DB.Collection("articulosMenu").
		DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artículo eliminado"})
}


// =======================
// DELETE MANY
// =======================
func DeleteManyArticulos(c *gin.Context) {

	_, err := config.DB.Collection("articulosMenu").
		DeleteMany(context.Background(), bson.M{"disponible": false})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Artículos eliminados"})
}


// =======================
// COUNT
// =======================
func CountArticulos(c *gin.Context) {

	count, err := config.DB.Collection("articulosMenu").
		CountDocuments(context.Background(), bson.M{})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"totalArticulos": count})
}


// =======================
// DISTINCT
// =======================
func DistinctCategorias(c *gin.Context) {

	result, err := config.DB.Collection("articulosMenu").
		Distinct(context.Background(), "categoria", bson.M{})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"categorias": result})
}