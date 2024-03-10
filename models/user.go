package models

import (
	"github.com/gocql/gocql"
)

type User struct {
	Id      gocql.UUID
	Name    string
	Surname string
	Email   string
}
