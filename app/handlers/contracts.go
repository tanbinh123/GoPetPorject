package handlers

import "github.com/Brigant/GoPetPorject/app/enteties"

type UserUsecase interface{
	CreateUser(enteties.User) (string, error)
}