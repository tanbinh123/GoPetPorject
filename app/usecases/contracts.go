package usecases

import "github.com/Brigant/GoPetPorject/app/enteties"

type Repository interface {
	AddUser(user enteties.User) (string, error)
}