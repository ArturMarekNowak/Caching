package controllers

import (
	"caching/internal/api/http/services"
	"caching/pkg/api/requests"
	"caching/pkg/api/responses"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		err := c.JSON(responses.NewHttpError("invalid path parameter"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(400)
	}
	user, err := services.GetUser(id)
	if err != nil {
		err := c.JSON(responses.NewHttpError("user not found"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(404)
	}
	err = c.JSON(user)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func CreateUser(c fiber.Ctx) error {
	var createUser requests.CreateUser
	err := c.Bind().Body(&createUser)
	if err != nil {
		err := c.JSON(responses.NewHttpError("invalid body"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(400)
	}
	id := services.CreateUser(createUser)
	err = c.JSON(responses.UserCreated{Id: id})
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}

func UpdateUser(c fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		err := c.JSON(responses.NewHttpError("invalid path parameter"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(400)
	}
	var createUser requests.CreateUser
	err = c.Bind().Body(&createUser)
	if err != nil {
		err := c.JSON(responses.NewHttpError("invalid body"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(400)
	}
	user, err := services.UpdateUser(id, createUser)
	if err != nil {
		err := c.JSON(responses.NewHttpError("user not found"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(404)
	}
	err = c.JSON(user)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func DeleteUser(c fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		err := c.JSON(responses.NewHttpError("invalid path parameter"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(400)
	}
	err = services.DeleteUser(id)
	if err != nil {
		err := c.JSON(responses.NewHttpError("user not found"))
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(404)
	}
	return c.SendStatus(204)
}
