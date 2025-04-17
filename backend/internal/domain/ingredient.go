package domain

import "database/sql"

type Ingredient struct {
	UUID         string       `json:"uuid"`
	Name         string       `json:"name"`
	CauseAllergy bool         `json:"cause_alergy"`
	Type         int          `json:"type"`
	Status       int          `json:"status"`
	CreatedAt    sql.NullTime `json:"createdAt"` // Changed to sql.NullTime
	UpdatedAt    sql.NullTime `json:"updatedAt"` // Changed to sql.NullTime
	DeletedAt    sql.NullTime `json:"deletedAt"` // Changed to sql.NullTime
}
