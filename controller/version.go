package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type Version struct {
	Name  string `json:"name"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
}

func VersionController(w http.ResponseWriter, r *http.Request) {
	version := Version{"first step", 0, 1}
	encoder := json.NewEncoder(w)

	err := encoder.Encode(&version)

	if err != nil {
		log.Printf("HTTP %s", err)
	}

}
