package handlers

import (
	"github.com/Brigant/GoPetPorject/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUsecase UserUsecase
	logger      *logger.Logger
}

func NewHandler(usecase UserUsecase, logger *logger.Logger) Handler {
	return Handler{
		userUsecase: usecase,
		logger:      logger,
	}
}

func (h *Handler) InitRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.Use(gin.Recovery())

	user := router.Group("/user")
	{
		user.POST("/sign-up", h.signUp)
		user.POST("/sign-in", h.signIn)
		user.GET("/logout", h.userIdentity, h.logout)
	}

	return router
}
