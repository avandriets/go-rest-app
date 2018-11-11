package controller

import (
	"encoding/json"
	"github.com/go-pg/pg/orm"
	"go-rest-app/model"
	"log"
	"net/http"
)

func StructuresByORM(w http.ResponseWriter, r *http.Request) {
	var structures []model.Structure
	encoder := json.NewEncoder(w)

	err := model.GetPgDatabase().Model(&structures).
		Apply(orm.Pagination(r.URL.Query())).Select()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := encoder.Encode(&structures); err != nil {
		log.Printf("HTTP %s", err)
	}
}

func StructuresController(w http.ResponseWriter, r *http.Request) {
	content := model.ContentResponse{}
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

func queryStructures(structures *model.ContentResponse) error {
	db := model.GetDatabase()

	rows, err := db.Query(`SELECT * FROM nprl.public.structures`)

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		structure := model.Structure{}
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
