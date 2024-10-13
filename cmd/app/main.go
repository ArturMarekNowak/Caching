package main

import (
	"caching/internal/api/http/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureEndpoints(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}

func ConfigureServer() {

	router := gin.Default()
	ConfigureEndpoints(router)
	router.Run(":8080")
}

func main() {
	ConfigureServer()
}
