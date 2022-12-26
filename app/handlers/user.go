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

	if err := c.BindJSON(&user); err != nil {
		h.logger.Infof("wrong body: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	h.logger.Debugf("request user: {Phone:%v, Age:%v}", user.Phone, user.Age)

	id, err := h.userUsecase.CreateUser(user)
	if err != nil {
		if errors.Is(err, enteties.ErrDuplicatePhone) {
			h.logger.Debugw("userUsecase", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		h.logger.Errorw("userUsecase", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {

}

func (h *Handler) logout(c *gin.Context) {

}
