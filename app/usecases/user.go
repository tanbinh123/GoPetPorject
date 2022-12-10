package usecases

import "github.com/Brigant/GoPetPorject/app/repositories/pg"

type Usecase struct {
	Authorization
}

type Authorization interface {
}

func NewUsecase(repo *pg.Repository) *Usecase {
	return &Usecase{}
}
