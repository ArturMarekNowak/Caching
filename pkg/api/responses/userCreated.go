package responses

import "github.com/gocql/gocql"

type UserCreated struct {
	Id gocql.UUID
}
