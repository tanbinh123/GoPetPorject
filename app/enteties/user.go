package enteties

import (
	"errors"
)

type User struct {
	ID       string `json:"id"`
	Phone    int    `json:"phone" binding:"gte=1,lte=999999999" `
	Password string `json:"password" binding:"min=8,max=255,ascii"`
	Age      int    `json:"age" binding:"gte=1,lte=150"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}

var (
	ErrDuplicatePhone = errors.New("phone is already exists in the database")
	ErrUserNotFound   = errors.New("user is not found with such credentials")
)
