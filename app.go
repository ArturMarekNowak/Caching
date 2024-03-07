package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"caching/controllers"
)

func ConfigureServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func ConfigureEndpoints() {

	http.HandleFunc("GET /users", controllers.GetUsers)
	http.HandleFunc("POST /users", controllers.CreateUser)
	http.HandleFunc("GET /users/{id}", controllers.GetUser)
	http.HandleFunc("PUT /users/{id}", controllers.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)
}

func main() {
	ConfigureEndpoints()
	ConfigureServer()
}
