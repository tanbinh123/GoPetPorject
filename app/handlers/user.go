package handlers

import (
	"errors"
	"net/http"

	"github.com/Brigant/GoPetPorject/app/enteties"
	"github.com/gin-gonic/gin"
)

// signUp func creates new user.
func (h *Handler) signUp(c *gin.Context) {
	var user enteties.User

	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Infof("wrong body: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	h.logger.Debugw("signUp", "phone", user.Phone, "age", user.Age)

	id, err := h.userUsecase.CreateUser(user)
	if err != nil {
		if errors.Is(err, enteties.ErrDuplicatePhone) {
			h.logger.Debugw("CreateUser", "error", err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Errorw("CreateUser", "error", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

type signInInput struct {
	Phone    int
	Password string
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		h.logger.Infof("wrong body: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	h.logger.Debugw("signIn", "phone", input.Phone, "password", input.Password)

	token, err := h.userUsecase.GenerateToken(input.Phone, input.Password)
	if err != nil {
		if errors.Is(err, enteties.ErrUserNotFound) {
			h.logger.Debugw("GenerateToken", "error", err.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

			return
		}

		h.logger.Errorw("GenerateToken", "error", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (h *Handler) logout(c *gin.Context) {

}
