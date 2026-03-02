package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ubicacion struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}

type Restaurante struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Nombre        string             `bson:"nombre"`
	Descripcion   string             `bson:"descripcion"`
	Categorias    []string           `bson:"categorias"`
	FechaRegistro time.Time          `bson:"fechaRegistro"`
	Estado        string             `bson:"estado"`
	Ubicacion     Ubicacion          `bson:"ubicacion"`
}