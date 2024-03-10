package services

import (
	"caching/models"
	"caching/repositories"
	"github.com/gocql/gocql"
)

func CreateUser(createUser models.CreateUser) gocql.UUID {
	return repositories.CreateUser(createUser)
}

func GetUser(id gocql.UUID) (*models.User, error) {
	user, err := repositories.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id gocql.UUID, updateUser models.CreateUser) (*models.User, error) {
	user, err := repositories.UpdateUser(id, updateUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id gocql.UUID) error {
	err := repositories.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
