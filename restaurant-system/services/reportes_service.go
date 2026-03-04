package services

import (
	"context"

	"restaurant-system/config"

	"go.mongodb.org/mongo-driver/bson"
)

func VentasPorMesService() ([]bson.M, error) {

	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id": bson.M{
					"restaurante": "$restaurante_id",
					"mes": bson.M{"$month": "$fecha"},
				},
				"totalVentas": bson.M{"$sum": "$total"},
			},
		},
		{
			"$sort": bson.M{
				"_id.mes": 1,
			},
		},
	}

	cursor, err := config.DB.Collection("ordenes").
		Aggregate(context.Background(), pipeline)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var resultados []bson.M
	cursor.All(context.Background(), &resultados)

	return resultados, nil
}