package controller

import (
	"database/sql"
	"encoding/json"
	"go-rest-app/model"
	"log"
	"net/http"
)

type Structure struct {
	NoID            int            `json:"noId"`
	Compounds       sql.NullString `json:"compounds"`
	Smiles          sql.NullString `json:"smiles"`
	Formular        sql.NullString `json:"formular"`
	Structure       sql.NullString `json:"structure"`
	StructureType   sql.NullString `json:"structure_type"`
	BioactivityType sql.NullString `json:"bioactivity_type"`
	Activity        sql.NullString `json:"activity"`
	JournalName     sql.NullString `json:"journalName"`
	Year            sql.NullInt64  `json:"year"`
	Volume          sql.NullInt64  `json:"volume"`
	Page            sql.NullInt64  `json:"page"`
	ArticleName     sql.NullString `json:"articleName"`
	ArticleNo       sql.NullInt64  `json:"articleNo"`
	CreatedAt       sql.NullString `json:"created_at"`
	UpdatedAt       sql.NullString `json:"updated_at"`
}

type ContentResponse struct {
	Structures []Structure `json:"content"`
}

func StructuresController(w http.ResponseWriter, r *http.Request) {
	content := ContentResponse{}
	encoder := json.NewEncoder(w)

	err := queryStructures(&content)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := encoder.Encode(&content); err != nil {
		log.Printf("HTTP %s", err)
	}
}

func queryStructures(structures *ContentResponse) error {
	db := model.GetDatabase()

	rows, err := db.Query(`SELECT * FROM nprl.public.structures`)

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		structure := Structure{}
		err = rows.Scan(
			&structure.NoID,
			&structure.Compounds,
			&structure.Smiles,
			&structure.Formular,
			&structure.Structure,
			&structure.StructureType,
			&structure.BioactivityType,
			&structure.Activity,
			&structure.JournalName,
			&structure.Year,
			&structure.Volume,
			&structure.Page,
			&structure.ArticleName,
			&structure.ArticleNo,
			&structure.CreatedAt,
			&structure.UpdatedAt,
		)
		if err != nil {
			return err
		}

		structures.Structures = append(structures.Structures, structure)
	}
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}
