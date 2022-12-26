package usecases

import (
	"crypto/sha256"
	"fmt"

	"github.com/Brigant/GoPetPorject/app/enteties"
)

type User struct {
	Repo Repository
}

func NewUser(repo Repository) User {
	return User{Repo: repo}
}

// The function CreateUser represents bissness logic layer 
// and  execute some manipulation with data before 
// saving  this data in the repository
func (u User) CreateUser(user enteties.User) (string, error) {
	user.Password = u.sha256(user.Password)

	id, err := u.Repo.AddUser(user)
	if err != nil {
		return "", fmt.Errorf("error occures while AddUser: %w",err)
	}

	return id, nil
}

// The function sha256 returns the checksum of the string using SHA256 hash algorithms.
func (u User) sha256(someString string) string {
	h := sha256.Sum256([]byte(someString))

	return fmt.Sprintf("%x", h)
}
