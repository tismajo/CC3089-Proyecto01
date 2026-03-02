package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ArticuloMenu struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Nombre        string             `bson:"nombre"`
	Descripcion   string             `bson:"descripcion"`
	Precio        float64            `bson:"precio"`
	Categoria     string             `bson:"categoria"`
	Disponible    bool               `bson:"disponible"`
	RestauranteID primitive.ObjectID `bson:"restaurante_id"`
}