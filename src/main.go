package main

import (
	"booleans/controllers"
	"booleans/services"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	services.Init()

	router.POST("/", controllers.MyController.NewBoolean)
	router.GET("/:id", controllers.MyController.GetBoolean)
	router.PATCH("/:id", controllers.MyController.UpdateBoolean)
	router.DELETE("/:id", controllers.MyController.DeleteBoolean)

	router.Run(":8080")
}
