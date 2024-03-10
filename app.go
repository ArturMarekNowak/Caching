package main

import (
	"caching/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints() {

	router := gin.Default()

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run(":8080")
}

func main() {
	ConfigureEndpoints()
}
