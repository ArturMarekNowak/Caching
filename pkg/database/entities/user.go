package entities

import (
	"encoding/json"
	"github.com/gocql/gocql"
)

type User struct {
	Id      gocql.UUID
	Name    string
	Surname string
	Email   string
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
