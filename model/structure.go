package model

import (
	"database/sql"
	"fmt"
)

type Structure struct {
	NoID            int            `json:"noId"sql:"noId"`
	Compounds       sql.NullString `json:"compounds"`
	Smiles          sql.NullString `json:"smiles"`
	Formular        sql.NullString `json:"formular"`
	Structure       sql.NullString `json:"structure"`
	StructureType   sql.NullString `json:"structure_type"`
	BioactivityType sql.NullString `json:"bioactivity_type"`
	Activity        sql.NullString `json:"activity"`
	JournalName     sql.NullString `json:"journalName"sql:"journalName"`
	Year            sql.NullInt64  `json:"year"`
	Volume          sql.NullInt64  `json:"volume"`
	Page            sql.NullInt64  `json:"page"`
	ArticleName     sql.NullString `json:"articleName"sql:"articleName"`
	ArticleNo       sql.NullInt64  `json:"articleNo"sql:"articleNo"`
	CreatedAt       sql.NullString `json:"created_at"`
	UpdatedAt       sql.NullString `json:"updated_at"`
}

func (u Structure) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.NoID, u.Smiles, u.Formular)
}

type ContentResponse struct {
	Structures []Structure `json:"content"`
}
