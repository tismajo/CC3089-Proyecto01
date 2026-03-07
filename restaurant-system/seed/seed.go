package seed

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"restaurant-system/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedBaseData() error {

	ctx := context.Background()

	usuariosCol := config.DB.Collection("usuarios")
	restCol := config.DB.Collection("restaurantes")
	artCol := config.DB.Collection("articulosMenu")

	rand.Seed(time.Now().UnixNano())

	// ===== USUARIOS =====
	var usuarios []interface{}
	for i := 0; i < 100; i++ {
		usuarios = append(usuarios, bson.M{
			"nombre":         fmt.Sprintf("Usuario %d", i),
			"correo":         fmt.Sprintf("usuario%d@mail.com", i),
			"contrasenaHash": "123456",
			"fechaRegistro":  time.Now(),
			"direccion":      "Ciudad",
			"roles":          []string{"cliente"},
		})
	}

	_, err := usuariosCol.InsertMany(ctx, usuarios)
	if err != nil {
		return err
	}

	// ===== RESTAURANTES =====
	var restaurantes []interface{}
	for i := 0; i < 50; i++ {
		restaurantes = append(restaurantes, bson.M{
			"nombre":        fmt.Sprintf("Restaurante %d", i),
			"descripcion":   "Comida deliciosa",
			"categorias":    []string{"categoria1", "categoria2"},
			"fechaRegistro": time.Now(),
			"estado":        "activo",
			"ubicacion": bson.M{
				"type":        "Point",
				"coordinates": []float64{-90.5 + rand.Float64(), 14.6 + rand.Float64()},
			},
		})
	}

	resultRest, err := restCol.InsertMany(ctx, restaurantes)
	if err != nil {
		return err
	}

	// ===== ARTICULOS =====
	var articulos []interface{}

	for _, id := range resultRest.InsertedIDs {

		restID := id.(primitive.ObjectID)

		for j := 0; j < 10; j++ {
			articulos = append(articulos, bson.M{
				"nombre":         fmt.Sprintf("Producto %d", j),
				"descripcion":    "Producto demo",
				"precio":         rand.Intn(100) + 20,
				"categoria":      "general",
				"disponible":     true,
				"restaurante_id": restID,
			})
		}
	}

	_, err = artCol.InsertMany(ctx, articulos)
	if err != nil {
		return err
	}

	fmt.Println("Seed base completado")
	return nil
}

func SeedOrdenes(n int) error {

	ctx := context.Background()

	usuariosCol := config.DB.Collection("usuarios")
	restCol := config.DB.Collection("restaurantes")
	ordenesCol := config.DB.Collection("ordenes")

	var usuarios []bson.M
	cursorU, _ := usuariosCol.Find(ctx, bson.M{})
	cursorU.All(ctx, &usuarios)

	var restaurantes []bson.M
	cursorR, _ := restCol.Find(ctx, bson.M{})
	cursorR.All(ctx, &restaurantes)

	rand.Seed(time.Now().UnixNano())

	batchSize := 1000

	for i := 0; i < n; i += batchSize {

		var docs []interface{}

		for j := 0; j < batchSize && i+j < n; j++ {

			usuario := usuarios[rand.Intn(len(usuarios))]
			restaurante := restaurantes[rand.Intn(len(restaurantes))]

			usuarioID := usuario["_id"].(primitive.ObjectID)
			restauranteID := restaurante["_id"].(primitive.ObjectID)

			cantidad := rand.Intn(3) + 1
			precio := float64(rand.Intn(100) + 20)
			subtotal := precio * float64(cantidad)

			item := bson.M{
				"articulo_id":     primitive.NewObjectID(),
				"nombre":          "Producto Demo",
				"precioUnitario":  precio,
				"cantidad":        cantidad,
				"subtotal":        subtotal,
			}

			fecha := time.Now().AddDate(0, -rand.Intn(12), -rand.Intn(30))

			doc := bson.M{
				"fecha":          fecha,
				"estado":         "completado",
				"total":          subtotal,
				"usuario_id":     usuarioID,
				"restaurante_id": restauranteID,
				"items":          []interface{}{item},
			}

			docs = append(docs, doc)
		}

		_, err := ordenesCol.InsertMany(ctx, docs)
		if err != nil {
			return err
		}

		fmt.Printf("Insertados %d documentos\n", i+batchSize)
	}

	fmt.Println("Seed ordenes completado")
	return nil
}