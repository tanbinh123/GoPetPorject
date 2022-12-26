package pg

import (
	"errors"
	"fmt"

	"github.com/Brigant/GoPetPorject/app/enteties"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const ErrCodeUniqueViolation = "unique_violation"

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{DB: db}
}

func (r Repository) AddUser(user enteties.User) (string, error) {
	var id string

	query := `INSERT INTO "user" (phone, password, age) VALUES ($1, $2, $3) RETURNING id`
	if err := r.DB.QueryRow(query, user.Phone, user.Password, user.Age).Scan(&id); err != nil {
		pqErr := new(pq.Error)
		if errors.As(err, &pqErr) && pqErr.Code.Name() == ErrCodeUniqueViolation {
			return "", enteties.ErrDuplicatePhone
		}

		return "", fmt.Errorf("cannot execute query: %w", err)
	}

	return id, nil
}
