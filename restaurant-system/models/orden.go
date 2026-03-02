package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemEmbedido struct {
	ArticuloID     primitive.ObjectID `bson:"articulo_id"`
	Nombre         string             `bson:"nombre"`
	PrecioUnitario float64            `bson:"precioUnitario"`
	Cantidad       int                `bson:"cantidad"`
	Subtotal       float64            `bson:"subtotal"`
}

type Orden struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Fecha         time.Time          `bson:"fecha"`
	Estado        string             `bson:"estado"`
	Total         float64            `bson:"total"`
	UsuarioID     primitive.ObjectID `bson:"usuario_id"`
	RestauranteID primitive.ObjectID `bson:"restaurante_id"`
	Items         []ItemEmbedido     `bson:"items"`
}