package services

import (
	"context"
	"time"

	"restaurant-system/config"
	"restaurant-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CrearOrdenConTransaccion(orden models.Orden) error {

	ctx := context.Background()

	session, err := config.Client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	return mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {

		if err := session.StartTransaction(); err != nil {
			return err
		}

		orden.Fecha = time.Now()

		_, err := config.DB.Collection("ordenes").InsertOne(sc, orden)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		_, err = config.DB.Collection("restaurantes").UpdateOne(
			sc,
			bson.M{"_id": orden.RestauranteID},
			bson.M{"$inc": bson.M{"ventasTotales": orden.Total}},
		)

		if err != nil {
			session.AbortTransaction(sc)
			return err
		}

		return session.CommitTransaction(sc)
	})
}