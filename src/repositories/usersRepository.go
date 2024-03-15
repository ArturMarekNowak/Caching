package repositories

import (
	"caching/src/models"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"log"
	"os"
)

func CreateUser(createUser models.CreateUser) gocql.UUID {

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
			log.Printf("insert public.users", err)
		}
	}()

	return uuid
}

func GetUser(id gocql.UUID) (*models.User, error) {

	var session = CreateSession()
	defer session.Close()
	stmt, names := qb.Select("public.users").Where(qb.Eq("Id")).ToCql()
	q := session.Query(stmt, names).BindMap(qb.M{
		"Id": gocql.UUID.String(id),
	})

	var user models.User
	if err := q.GetRelease(&user); err != nil {
		log.Printf("select public.users", err)
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id gocql.UUID, createUser models.CreateUser) (*models.User, error) {

	if _, err := GetUser(id); err != nil {
		log.Printf("select public.users", err)
		return nil, err
	}

	var session = CreateSession()
	defer session.Close()
	stmt, names := qb.Update("public.users").Set("Name", "Surname", "Email").Where(qb.Eq("Id")).ToCql()
	session.Query(stmt, names).BindMap(qb.M{
		"Name":    createUser.Name,
		"Surname": createUser.Surname,
		"Email":   createUser.Email,
		"Id":      gocql.UUID.String(id),
	}).Exec()

	return &models.User{Id: id, Name: createUser.Name, Surname: createUser.Surname, Email: createUser.Email}, nil
}

func DeleteUser(id gocql.UUID) error {

	if _, err := GetUser(id); err != nil {
		log.Printf("delete public.users", err)
		return err
	}

	var session = CreateSession()
	defer session.Close()
	stmt, names := qb.Delete("public.users").Where(qb.Eq("Id")).ToCql()
	session.Query(stmt, names).BindMap(qb.M{
		"Id": gocql.UUID.String(id),
	}).Exec()

	return nil
}

func CreateSession() gocqlx.Session {
	var cluster = gocql.NewCluster(os.Getenv("CONNECTION_STRING"))
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Printf("unable to connect to cassandra", err)
	}
	return session
}
