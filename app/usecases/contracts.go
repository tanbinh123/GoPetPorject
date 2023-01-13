package usecases

import (
	"time"

	"github.com/Brigant/GoPetPorject/app/enteties"
)

type Repository interface {
	AddUser(user enteties.User) (string, error)
	GetUser(phone int, password string) (string, error)
	AddSession(userID string, refreshTokenTTL time.Duration) (string, error)
	DeleteSession(token string) error
	GetSession(refreshToken string) (enteties.Session, error)
	UpdateSession(refreshToken string, refreshTokenTTL time.Duration) error
	
}
