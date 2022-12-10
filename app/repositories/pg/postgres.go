package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("cannot connet to db: s%", err)
	}

	return db, nil
}
