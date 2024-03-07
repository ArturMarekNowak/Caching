package controllers

import (
	"caching/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, World!")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))
	user := repositories.GetUser(id)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
