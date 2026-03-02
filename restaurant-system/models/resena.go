package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resena struct {
	ID            primitive.ObjectID  `bson:"_id,omitempty"`
	Calificacion  int                 `bson:"calificacion"`
	Comentario    string              `bson:"comentario"`
	Fecha         time.Time           `bson:"fecha"`
	UsuarioID     primitive.ObjectID  `bson:"usuario_id"`
	RestauranteID primitive.ObjectID  `bson:"restaurante_id"`
	OrdenID       *primitive.ObjectID `bson:"orden_id,omitempty"`
}