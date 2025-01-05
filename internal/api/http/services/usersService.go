package services

import (
	"caching/internal/api/http/repositories"
	"caching/internal/helpers/caching"
	"caching/pkg/api/requests"
	models "caching/pkg/database/entities"
	"github.com/gocql/gocql"
	"log"
)

func CreateUser(createUser requests.CreateUser) gocql.UUID {
	return repositories.CreateUser(createUser)
}

func GetUser(id gocql.UUID) (*models.User, error) {
	var cachedUser models.User
	err := caching.GetKey(id.String(), &cachedUser)
	if err == nil {
		log.Print("Cache hit")
		return &cachedUser, nil
	}
	log.Print("Cache miss")
	user, err := repositories.GetUser(id)
	if err != nil {
		return nil, err
	}
	err = caching.SetKey(id.String(), user)
	if err != nil {
		log.Print("Couldn't save key")
	}
	return user, nil
}

func UpdateUser(id gocql.UUID, updateUser requests.CreateUser) (*models.User, error) {
	user, err := repositories.UpdateUser(id, updateUser)
	if err != nil {
		return nil, err
	}
	var cachedUser models.User
	if caching.KeyExists(id.String(), cachedUser) {
		err = caching.SetKey(id.String(), user)
		if err != nil {
			log.Print("Couldn't save key")
		}
	}
	return user, nil
}

func DeleteUser(id gocql.UUID) error {
	err := repositories.DeleteUser(id)
	if err != nil {
		return err
	}
	err = caching.DelKey(id.String())
	if err != nil {
		log.Print("Couldn't del key")
	}
	return nil
}
