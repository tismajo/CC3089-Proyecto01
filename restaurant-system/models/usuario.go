package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre         string             `bson:"nombre"`
	Correo         string             `bson:"correo"`
	ContrasenaHash string             `bson:"contrasenaHash"`
	FechaRegistro  time.Time          `bson:"fechaRegistro"`
	Direccion      string             `bson:"direccion"`
	Roles          []string           `bson:"roles"`
}