package handlers

import (
	"errors"
	"net/http"

	"github.com/Brigant/GoPetPorject/app/enteties"
	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Phone    int
	Password string
}

type inputRefreshToken struct {
	RefreshToken string
}

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

	id, err := h.authUsecase.CreateUser(user)
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

// Respond with accessToken and refreshToken if user
// provides the valid credentials.
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Infof("wrong body: %v", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	h.logger.Debugw("signIn", "phone", input.Phone)

	accessToken, refreshToken, err := h.authUsecase.GenerateToken(input.Phone, input.Password)
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

	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

// Recreate new accessToken and new refreshToken if refreshToken is valid.
func (h *Handler) refreshTokens(c *gin.Context) {
	var rToken inputRefreshToken

	if err := c.BindJSON(&rToken); err != nil {
		h.logger.Infof("wrong body: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	accesToken, refrestToken, err := h.authUsecase.RefreshTokens(rToken.RefreshToken)
	if err != nil {
		if errors.Is(err, enteties.ErrRefreshTokenExpired) || errors.Is(err, enteties.ErrSesseionNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": accesToken, "refreshToken": refrestToken})
}

// Make logout.
func (h *Handler) logout(c *gin.Context) {
	var rToken inputRefreshToken

	id, _ := c.Get(userCtx)

	if err := c.BindJSON(&rToken); err != nil {
		h.logger.Infof("wrong body: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := h.authUsecase.DeleteToken(rToken.RefreshToken); err != nil {
		h.logger.Errorf("wrong body: %v", err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId": id,
	})
}
