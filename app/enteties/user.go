package enteties

import "errors"

type User struct {
	ID       string `json:"id"`
	Phone    int    `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

var (
	ErrDuplicatePhone = errors.New("phone is already exists in the database")
	ErrUserNotFound = errors.New("user is not found with such credentials")
)

