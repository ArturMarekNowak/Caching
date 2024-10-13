package controllers

import (
	"caching/internal/api/http/services"
	"caching/internal/helpers"
	"caching/pkg/api/requests"
	"caching/pkg/api/responses"
	"caching/pkg/database/entities"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
	"log"
)

func GetUser(c fiber.Ctx) error {

	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		c.JSON(responses.NewHttpError("invalid path parameter"))
		return c.SendStatus(400)
	}

	var cachedUser entities.User
	err = helpers.GetKey(id.String(), &cachedUser)
	if err == nil {
		c.JSON(cachedUser)
		log.Print("Cache hit")
		return c.SendStatus(200)
	}

	log.Print("Cache miss")

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(responses.NewHttpError("user not found"))
		return c.SendStatus(404)
	}
	c.JSON(user)

	err = helpers.SetKey(id.String(), user)
	if err != nil {
		log.Print("Could save key %s", id)
	}
	return c.SendStatus(200)
}

func CreateUser(c fiber.Ctx) error {
	var createUser requests.CreateUser
	err := c.Bind().Body(&createUser)
	if err != nil {
		c.JSON(responses.NewHttpError("invalid body"))
		return c.SendStatus(400)
	}
	id := services.CreateUser(createUser)
	c.JSON(responses.UserCreated{Id: id})

	return c.SendStatus(201)
}

func UpdateUser(c fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		c.JSON(responses.NewHttpError("invalid path parameter"))
		return c.SendStatus(400)
	}
	var createUser requests.CreateUser
	err = c.Bind().Body(&createUser)
	if err != nil {
		c.JSON(responses.NewHttpError("invalid body"))
		return c.SendStatus(400)
	}
	user, err := services.UpdateUser(id, createUser)
	if err != nil {
		c.JSON(responses.NewHttpError("user not found"))
		return c.SendStatus(404)
	}
	c.JSON(user)

	var cachedUser entities.User
	if helpers.GetKey(id.String(), &cachedUser) == nil {
		err = helpers.SetKey(id.String(), user)
		if err != nil {

			log.Print("Couldn't save key %s", id)
		}
	}

	return c.SendStatus(200)
}

func DeleteUser(c fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		c.JSON(responses.NewHttpError("invalid path parameter"))
		return c.SendStatus(400)
	}
	err = services.DeleteUser(id)
	if err != nil {
		c.JSON(responses.NewHttpError("user not found"))
		return c.SendStatus(404)
	}

	err = helpers.DelKey(id.String())
	if err != nil {
		log.Print("Couldn't del key %s", id)
	}
	return c.SendStatus(204)
}
