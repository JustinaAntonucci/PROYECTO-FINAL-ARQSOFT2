package router

import (
	"fmt"
	hotelController "hotels-api/controllers/hotel"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {

	// router.Use(corsMiddleware())
	// Products Mapping
	router.GET("/hotels-api/hotels/:id", hotelController.Get)
	router.POST("/hotels-api/hotels", hotelController.Insert)
	router.PUT("/hotels-api/hotels/:id", hotelController.Update)
	router.DELETE("/hotels-api/hotels/:id", hotelController.Delete) // Ruta para eliminar un hotel por ID
	router.OPTIONS("/hotels-api/hotels", hotelController.Insert)
	router.OPTIONS("/hotels-api/hotels/:id", hotelController.Update)

	fmt.Println("Finishing mappings configurations")
}

func corsMiddleware() gin.HandlerFunc {
	// Define your CORS configuration here
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your frontend's origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	return cors.New(config)
}
