package handlers

import (
	"github.com/Brigant/GoPetPorject/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authUsecase AuthUsecase
	logger      *logger.Logger
}

func NewHandler(usecase AuthUsecase, logger *logger.Logger) Handler {
	return Handler{
		authUsecase: usecase,
		logger:      logger,
	}
}

func (h *Handler) InitRouter(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.Use(gin.Recovery())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refreshTokens)
		auth.GET("/logout", h.userIdentity, h.logout)
	}

	return router
}
