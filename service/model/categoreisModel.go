package model

import (
	"database/sql"
)

type Categoreis struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func GetCategoreisByID(db *sql.DB, Name string) error {
	sql := "INSERT INTO categoreis (name) VALUES (?)"
	_, err := db.Query(sql, Name)

	if err != nil {
		return err
	}
	return nil
}
