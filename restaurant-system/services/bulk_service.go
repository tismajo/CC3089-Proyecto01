package services

import (
	"context"
	"fmt"
	"time"

	"restaurant-system/config"
	"restaurant-system/models"

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
