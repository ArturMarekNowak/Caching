package repositories

import (
	"caching/pkg/api/requests"
	"caching/pkg/database/entities"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"log"
	"os"
)

func CreateUser(createUser requests.CreateUser) gocql.UUID {
	var session = CreateSession()
	uuid, _ := gocql.RandomUUID()
	stmt, names := qb.Insert("public.users").Columns("Id", "Name", "Surname", "Email").ToCql()
	q := session.Query(stmt, names).BindMap(qb.M{
		"Id":      gocql.UUID.String(uuid),
		"Name":    createUser.Name,
		"Surname": createUser.Surname,
		"Email":   createUser.Email,
	})
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			log.Printf("insert public.users %s", err)
		}
	}()
	return uuid
}

func GetUser(id gocql.UUID) (*entities.User, error) {
	var session = CreateSession()
	defer session.Close()
	stmt, names := qb.Select("public.users").Where(qb.Eq("Id")).ToCql()
	q := session.Query(stmt, names).BindMap(qb.M{
		"Id": gocql.UUID.String(id),
	})
	var user entities.User
	if err := q.GetRelease(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id gocql.UUID, createUser requests.CreateUser) (*entities.User, error) {
	if _, err := GetUser(id); err != nil {
		return nil, err
	}
	var session = CreateSession()
	defer session.Close()
	stmt, names := qb.Update("public.users").Set("Name", "Surname", "Email").Where(qb.Eq("Id")).ToCql()
	err := session.Query(stmt, names).BindMap(qb.M{
		"Name":    createUser.Name,
		"Surname": createUser.Surname,
		"Email":   createUser.Email,
		"Id":      gocql.UUID.String(id),
	}).Exec()
	if err != nil {
		return nil, err
	}
	return &entities.User{Id: id, Name: createUser.Name, Surname: createUser.Surname, Email: createUser.Email}, nil
}

func DeleteUser(id gocql.UUID) error {
	if _, err := GetUser(id); err != nil {
		log.Printf("delete public.users %s", err)
		return err
	}
	var session = CreateSession()
	defer session.Close()
	stmt, names := qb.Delete("public.users").Where(qb.Eq("Id")).ToCql()
	err := session.Query(stmt, names).BindMap(qb.M{
		"Id": gocql.UUID.String(id),
	}).Exec()
	if err != nil {
		return err
	}
	return nil
}

func CreateSession() gocqlx.Session {
	var cluster = gocql.NewCluster(os.Getenv("CONNECTION_STRING"))
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Printf("unable to connect to cassandra %s", err)
	}
	return session
}
