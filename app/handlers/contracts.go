package handlers

import "github.com/Brigant/GoPetPorject/app/enteties"

type AuthUsecase interface {
	CreateUser(enteties.User) (string, error)
	GenerateToken(int, string) (accessToken, refreshToken string, err error)
	RefreshTokens(string) (accessToken, refreshToken string, err error)
	ParseToken(token string) (string, error)
	DeleteToken(token string) error
}
