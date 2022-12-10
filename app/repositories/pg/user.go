package pg

type Repository struct {
	Authorization
}

type Authorization interface {
}

func NewRepository() *Repository {
	return &Repository{}
}
