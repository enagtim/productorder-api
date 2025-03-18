package db

import (
	"database/sql"
	"order-api/configs"

	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func NewDatabase(conf *configs.Config) *Db {
	db, err := sql.Open("postgres", conf.Db.Dsn)
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
