package pg

import "github.com/jmoiron/sqlx"

type Repository struct {
	Authorization
}

type Authorization interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
