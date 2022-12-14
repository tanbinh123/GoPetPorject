package handlers

import (
	"github.com/Brigant/GoPetPorject/app/usecases"
	"github.com/Brigant/GoPetPorject/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecases.Usecase
	logger *logger.Logger
}

func NewHandler(usecase *usecases.Usecase, logger *logger.Logger) *Handler {
	return &Handler{
		usecase: usecase,
		logger: logger,
	}
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
