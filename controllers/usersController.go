package controllers

import (
	"caching/helpers"
	"caching/models"
	"caching/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

func GetUser(c *gin.Context) {

	id, err := gocql.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}

	var cachedUser models.User
	err = helpers.GetKey(id.String(), &cachedUser)
	if err == nil {
		c.JSON(http.StatusOK, cachedUser)
		log.Print("Cache hit")
		return
	}

	fmt.Println("Cache miss")

	user, err := services.GetUser(id)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	c.JSON(http.StatusOK, user)

	err = helpers.SetKey(id.String(), user)
	if err != nil {
		fmt.Println("Could save key %s", id)
	}
}

func CreateUser(c *gin.Context) {
	var createUser models.CreateUser
	err := c.BindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	id := services.CreateUser(createUser)
	c.JSON(http.StatusCreated, models.UserCreated{Id: id})

	err = helpers.SetKey(id.String(), models.User{
		Id:      id,
		Name:    createUser.Name,
		Surname: createUser.Surname,
		Email:   createUser.Email})
	if err != nil {
		fmt.Println("Couldn't save key %s", id)
	}
}

func UpdateUser(c *gin.Context) {
	id, err := gocql.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	var createUser models.CreateUser
	err = c.BindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	user, err := services.UpdateUser(id, createUser)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	c.JSON(http.StatusOK, user)

	err = helpers.SetKey(id.String(), user)
	if err != nil {
		fmt.Println("Couldn't save key %s", id)
	}
}

func DeleteUser(c *gin.Context) {
	id, err := gocql.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	err = services.DeleteUser(id)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	c.JSON(http.StatusNoContent, nil)

	err = helpers.DelKey(id.String())
	if err != nil {
		fmt.Println("Couldn't del key %s", id)
	}
}
