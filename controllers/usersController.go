package controllers

import (
	"caching/models"
	"caching/services"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"net/http"
)

func GetUser(c *gin.Context) {

	id, err := gocql.ParseUUID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	user, err := services.GetUser(id)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var createUser models.CreateUser
	err := c.BindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid path parameter"})
	}
	user := services.CreateUser(createUser)
	c.JSON(http.StatusCreated, user)
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
}
