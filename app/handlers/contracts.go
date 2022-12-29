package handlers

import "github.com/Brigant/GoPetPorject/app/enteties"

type UserUsecase interface{
	CreateUser(enteties.User) (string, error)
	GenerateToken(int, string) (string, error)
}