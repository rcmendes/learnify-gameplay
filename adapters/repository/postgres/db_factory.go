package postgres

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //Postgres driver
)

// Connect connects to the postgres database.
func Connect() *pg.DB {
	//TODO import from env or receive a config struct as parameter.
	return pg.Connect(&pg.Options{
		User:     "dev",
		Password: "dev",
		Database: "learnify_dev",
	})
}

func connect() *sqlx.DB {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=dev password=dev dbname=learnify_dev sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
