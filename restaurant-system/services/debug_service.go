package services

import (
	"context"

	"restaurant-system/config"

	"go.mongodb.org/mongo-driver/bson"
)

func ExplainOrdenesService() (bson.M, error) {

	command := bson.D{
		{"explain", bson.D{
			{"find", "ordenes"},
			{"filter", bson.D{{"estado", "completado"}}},
		}},
		{"verbosity", "executionStats"},
	}

	var result bson.M

	err := config.DB.RunCommand(
		context.Background(),
		command,
	).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}