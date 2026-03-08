package services

import (
	"context"
	"fmt"
	"time"

	"restaurant-system/config"
	"restaurant-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func BulkInsertUsuarios() error {

	var usuarios []interface{}

	// Creamos 5 usuarios dinámicos
	for i := 1; i <= 5; i++ {

		usuario := models.Usuario{
			Nombre:         fmt.Sprintf("BulkUser%d", i),
			Correo:         fmt.Sprintf("bulk%d_%d@mail.com", i, time.Now().Unix()),
			ContrasenaHash: "123456",
			Direccion:      fmt.Sprintf("Zona %d", i),
			Roles:          []string{"cliente"},
			FechaRegistro:  time.Now(),
		}

		usuarios = append(usuarios, usuario)
	}

	_, err := config.DB.Collection("usuarios").
		InsertMany(context.Background(), usuarios)

	return err
}

func BulkOperacionMixta() error {

	ctx := context.Background()

	var modelsWrite []mongo.WriteModel

	// ======================
	// INSERT
	// ======================
	for i := 1; i <= 2; i++ {

		usuario := models.Usuario{
			Nombre:         fmt.Sprintf("MixUser%d", i),
			Correo:         fmt.Sprintf("mix%d_%d@mail.com", i, time.Now().UnixNano()),
			ContrasenaHash: "123456",
			Direccion:      "Zona Bulk",
			Roles:          []string{"cliente"},
			FechaRegistro:  time.Now(),
		}

		insertModel := mongo.NewInsertOneModel().
			SetDocument(usuario)

		modelsWrite = append(modelsWrite, insertModel)
	}

	// ======================
	// UPDATE MANY
	// ======================
	updateModel := mongo.NewUpdateManyModel().
		SetFilter(bson.M{"roles": "cliente"}).
		SetUpdate(bson.M{
			"$set": bson.M{"direccion": "Actualizado por Bulk"},
		})

	modelsWrite = append(modelsWrite, updateModel)

	// ======================
	// DELETE ONE
	// ======================
	deleteModel := mongo.NewDeleteOneModel().
		SetFilter(bson.M{"nombre": "Usuario 0"})

	modelsWrite = append(modelsWrite, deleteModel)

	// ======================
	// EJECUTAR BULK
	// ======================
	_, err := config.DB.Collection("usuarios").
		BulkWrite(ctx, modelsWrite)

	return err
}