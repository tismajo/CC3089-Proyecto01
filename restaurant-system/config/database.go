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
	
	createValidations()
	createIndexes()

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
	// Índice para estado (necesario para explain)
	DB.Collection("ordenes").Indexes().CreateOne(context.Background(),
	mongo.IndexModel{
		Keys: bson.D{{"estado", 1}},
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

	ctx := context.Background()

	// =======================
	// USUARIOS
	// =======================
	usuarioSchema := bson.M{
		"bsonType": "object",
		"required": []string{"nombre", "correo"},
		"properties": bson.M{
			"nombre": bson.M{"bsonType": "string"},
			"correo": bson.M{"bsonType": "string"},
			"roles": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
		},
	}

	createCollectionWithValidation("usuarios", usuarioSchema, ctx)

	// =======================
	// RESTAURANTES
	// =======================
	restauranteSchema := bson.M{
		"bsonType": "object",
		"required": []string{"nombre", "estado", "ubicacion"},
		"properties": bson.M{
			"nombre": bson.M{"bsonType": "string"},
			"estado": bson.M{"bsonType": "string"},
			"categorias": bson.M{
				"bsonType": "array",
				"items":    bson.M{"bsonType": "string"},
			},
			"ubicacion": bson.M{
				"bsonType": "object",
				"required": []string{"type", "coordinates"},
				"properties": bson.M{
					"type": bson.M{"enum": []string{"Point"}},
					"coordinates": bson.M{
						"bsonType": "array",
						"items":    bson.M{"bsonType": "double"},
					},
				},
			},
		},
	}

	createCollectionWithValidation("restaurantes", restauranteSchema, ctx)

	// =======================
	// ARTICULOS
	// =======================
	articuloSchema := bson.M{
		"bsonType": "object",
		"required": []string{"nombre", "precio", "restaurante_id"},
		"properties": bson.M{
			"nombre": bson.M{"bsonType": "string"},
			"precio": bson.M{"bsonType": "double"},
			"restaurante_id": bson.M{
				"bsonType": "objectId",
			},
		},
	}

	createCollectionWithValidation("articulosMenu", articuloSchema, ctx)

	// =======================
	// ORDENES
	// =======================
	ordenSchema := bson.M{
		"bsonType": "object",
		"required": []string{"usuario_id", "restaurante_id", "estado", "items"},
		"properties": bson.M{
			"usuario_id": bson.M{"bsonType": "objectId"},
			"restaurante_id": bson.M{"bsonType": "objectId"},
			"estado": bson.M{
				"enum": []string{"pendiente", "procesando", "cancelado", "completado"},
			},
			"items": bson.M{
				"bsonType": "array",
			},
		},
	}

	createCollectionWithValidation("ordenes", ordenSchema, ctx)

	// =======================
	// RESEÑAS
	// =======================
	resenaSchema := bson.M{
		"bsonType": "object",
		"required": []string{"calificacion", "usuario_id", "restaurante_id"},
		"properties": bson.M{
			"calificacion": bson.M{
				"bsonType": "int",
				"minimum":  1,
				"maximum":  5,
			},
			"usuario_id": bson.M{"bsonType": "objectId"},
			"restaurante_id": bson.M{"bsonType": "objectId"},
		},
	}

	createCollectionWithValidation("resenas", resenaSchema, ctx)
}

func createCollectionWithValidation(name string, schema bson.M, ctx context.Context) {

	opts := options.CreateCollection().
		SetValidator(bson.M{"$jsonSchema": schema}).
		SetValidationLevel("strict")

	err := DB.CreateCollection(ctx, name, opts)

	if err != nil {
		log.Println(name, "ya existe o error:", err)
	}
}