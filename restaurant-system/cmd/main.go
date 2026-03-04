package main

import (
	"restaurant-system/config"
	"restaurant-system/handlers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {

	config.ConnectDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/usuarios", handlers.CreateUsuario)
	r.GET("/usuarios", handlers.GetUsuarios)
	r.POST("/usuarios/bulk", handlers.BulkUsuarios)

	r.POST("/restaurantes", handlers.CreateRestaurante)
	r.GET("/restaurantes", handlers.GetRestaurantes)
	r.GET("/restaurantes/cercanos", handlers.RestaurantesCercanos)

	r.POST("/ordenes", handlers.CreateOrden)
	r.GET("/ordenes", handlers.GetOrdenes)
	r.PUT("/ordenes/:id/cancelar", handlers.UpdateOrdenEstado)
	r.PUT("/ordenes/masivo", handlers.UpdateManyOrdenes)
	r.DELETE("/ordenes/:id", handlers.DeleteOrden)

	r.POST("/resenas", handlers.CreateResena)
	r.GET("/resenas", handlers.GetResenas)

	r.GET("/reportes/mejores-restaurantes", handlers.RestaurantesMejorCalificados)
	r.GET("/reportes/ventas-por-mes", handlers.VentasPorMes)
	r.GET("/reportes/platillos-mas-vendidos", handlers.PlatillosMasVendidos)

	r.GET("/debug/explain-ordenes", handlers.ExplainOrdenes)

	r.Run(":8080")
}