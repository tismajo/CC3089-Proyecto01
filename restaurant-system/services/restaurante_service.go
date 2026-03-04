package services

import (
	"context"

	"restaurant-system/config"
	"restaurant-system/models"

	"go.mongodb.org/mongo-driver/bson"
)

func RestaurantesCercanosService(lat, lng float64, maxDistance int) ([]models.Restaurante, error) {

	filter := bson.M{
		"ubicacion": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
				"$maxDistance": maxDistance,
			},
		},
	}

	cursor, err := config.DB.Collection("restaurantes").
		Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var restaurantes []models.Restaurante
	cursor.All(context.Background(), &restaurantes)

	return restaurantes, nil
}