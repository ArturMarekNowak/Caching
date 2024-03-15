package models

import "github.com/gocql/gocql"

type UserCreated struct {
	Id gocql.UUID
}
