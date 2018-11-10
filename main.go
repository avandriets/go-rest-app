package main

import (
	. "./controller"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-app/model"
	"log"
	"net/http"
)
import _ "github.com/lib/pq"
import "database/sql"

func main() {

	db := connectToDatabase()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/structures", StructuresController).Methods("GET")
	r.HandleFunc("/version", VersionController).Methods("GET")
	r.HandleFunc("/metadata", MetadataController).Methods("GET")

	if err := http.ListenAndServe(":9000", r); err != nil {
		log.Fatal(err)
	}

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
