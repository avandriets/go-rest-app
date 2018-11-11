package main

import (
	. "./controller"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-app/model"
	"log"
	"net/http"

	"github.com/go-pg/pg"
)
import _ "github.com/lib/pq"
import "database/sql"

func main() {
	db := connectToPgDataBase()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/structures", StructuresByORM).Methods("GET")
	r.HandleFunc("/version", VersionController).Methods("GET")
	r.HandleFunc("/metadata", MetadataController).Methods("GET")
	r.HandleFunc("/render", RenderStructureController).Methods("GET")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}
}

func connectToPgDataBase() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "nprl",
		Password: "nprl",
		Database: "nprl",
	})

	model.SetPgDataBase(db)
	return db
}

func connectToDatabase() *sql.DB {
	// TODO put connection parameters to environment variables
	connStr := "postgresql://nprl:nprl@localhost:5432/nprl?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable connect to DB: %v", err))
	}

	model.SetDatabase(db)

	return db
}
