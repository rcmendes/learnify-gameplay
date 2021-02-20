package postgres

import (
	"github.com/go-pg/pg/v10"
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
