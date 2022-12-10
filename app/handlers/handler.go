package handlers

import (
	"github.com/Brigant/GoPetPorject/app/usecases"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecases.Usecase
}

func NewHandler(usecase *usecases.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) InitRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)
		user.GET("/logout", h.logout)
	}

	return router
}
