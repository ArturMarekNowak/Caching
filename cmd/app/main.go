package main

import (
	"caching/internal/api/http/controllers"
	"github.com/gofiber/fiber/v3"
	"log"
)

func ConfigureEndpoints(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:id", controllers.GetUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}

func ConfigureServer() {
	app := fiber.New()
	ConfigureEndpoints(app)
	log.Fatal(app.Listen(":8080"))
}

func main() {
	ConfigureServer()
}
