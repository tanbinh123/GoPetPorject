package pg

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Brigant/GoPetPorject/app/enteties"
)

func (r Repository) AddSession(userID string, ttl time.Duration) (string, error) {
	var refreshToken string

	query := `INSERT INTO "session"
		(user_id, expires_in)
		VALUES($1, $2) RETURNING id`

	fmt.Println(ttl)

	if err := r.DB.QueryRow(query, userID, ttl).Scan(&refreshToken); err != nil {
		return "", fmt.Errorf("cannot execute query: %w", err)
	}

	return refreshToken, nil
}

func (r Repository) GetSession(refreshToken string) (enteties.Session, error) {
	var session enteties.Session

	query := `SELECT id, user_id, expires_in, created FROM "session" WHERE id=$1`
	row := r.DB.QueryRow(query, refreshToken)

	if err := row.Scan(&session.ID, &session.UserID, &session.ExpiresIn, &session.Created); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return  enteties.Session{}, enteties.ErrSesseionNotFound
		}

		return enteties.Session{}, err
	}

	return session, nil
}

func (r Repository) UpdateSession(refreshToken string, TTL time.Duration) error {
	query := `	UPDATE "session" 
				SET expires_in=$2
				WHERE id=$1;`

	if _, err := r.DB.Exec(query, refreshToken, TTL); err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteSession(refreshToken string) error {
	query := `DELETE FROM "session" WHERE id=$1;`
	if _, err := r.DB.Exec(query, refreshToken); err != nil {
		return err
	}

	return nil
}
