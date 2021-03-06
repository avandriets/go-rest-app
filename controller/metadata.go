package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Metadata struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	Display    bool   `json:"display"`
	Filterable bool   `json:"filterable"`
	Sortable   bool   `json:"sortable"`
	Type       string `json:"type"`
	Formatter  string `json:"formatter"`
	Link       string `json:"link"`
	FilterType string `json:"filterType"`
	Disabled   bool   `json:"disabled"`
}

func MetadataController(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("./static/metadata.json")

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Server error cannot get metadata", http.StatusInternalServerError)
	}
	fmt.Println("Successfully Opened metadata.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []Metadata
	json.Unmarshal([]byte(byteValue), &result)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(&result); err != nil {
		log.Printf("HTTP %s", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
