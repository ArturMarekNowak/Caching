package repositories

import (
	"caching/models"
	"github.com/gocql/gocql"
	"log"
)

func GetUser(id int) models.User {

	var c = gocql.NewCluster("host.docker.internal:9042")
	session, err := gocql.NewSession(*c)
	if err != nil {
		log.Fatal("unable to connect to cassandra")
	}
	defer session.Close()

	q := session.Query("SELECT name, surname, email FROM public.users WHERE Id = ?", id)

	var name, surname, email string

	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			log.Printf("select public.users", err)
		}
	}()

	for it.Scan(&name, &surname, &email) {
		log.Printf("\t" + name + " " + surname + ", " + email)
	}

	var user models.User

	user.Name = name
	user.Surname = surname
	user.Email = email

	return user
}
