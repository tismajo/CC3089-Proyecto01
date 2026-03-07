package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:1234@lab01.2h9pt9n.mongodb.net/?appName=lab01")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	DB = client.Database("restaurant_system")

	createIndexes()
	createValidations()

	log.Println("✅ Conectado a MongoDB")
}

func createIndexes() {

	// Índice simple unique
	DB.Collection("usuarios").Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{"correo", 1}},
			Options: options.Index().SetUnique(true),
		})

	// Índice compuesto
	DB.Collection("ordenes").Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{"usuario_id", 1}, {"fecha", -1}},
		})

	// Índice multikey
	DB.Collection("restaurantes").Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{"categorias", 1}},
		})

	// Índice geoespacial
	DB.Collection("restaurantes").Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{"ubicacion", "2dsphere"}},
		})

	// Índice de texto
	DB.Collection("articulosMenu").Indexes().CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{"nombre", "text"}, {"descripcion", "text"}},
		})
}

func createValidations() {

	usuarioSchema := bson.M{
		"bsonType": "object",
		"required": []string{"nombre", "correo"},
		"properties": bson.M{
			"nombre": bson.M{
				"bsonType": "string",
			},
			"correo": bson.M{
				"bsonType": "string",
			},
			"roles": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "string",
				},
			},
		},
	}

	command := bson.D{
		{"collMod", "usuarios"},
		{"validator", bson.M{"$jsonSchema": usuarioSchema}},
		{"validationLevel", "strict"},
	}

	err := DB.RunCommand(context.Background(), command).Err()
	if err != nil {
		log.Println("⚠ Validación no aplicada (puede que la colección no exista aún)")
	}
}