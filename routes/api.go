package routes

import (
	"project-golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routing Group biar rapi mirip Laravel
	v1 := r.Group("/products")
	{
		v1.GET("", controllers.FindProducts)
		v1.GET("/:id", controllers.FindProductById)
		v1.POST("", controllers.CreateProduct)
	}

	return r
}