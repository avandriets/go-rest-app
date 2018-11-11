package model

import "github.com/go-pg/pg"

var dbPg *pg.DB

func SetPgDataBase(database *pg.DB) {
	dbPg = database
}

func GetPgDatabase() *pg.DB {
	return dbPg
}
