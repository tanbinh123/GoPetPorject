package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	athoriazahionHeader = "Authorization"
	userCtx             = "userID"
)

var (
	errorEmptyHeader   = errors.New("empty header")
	errorInvalidHeader = errors.New("invalid header")
)

func (h Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(athoriazahionHeader)
	if header == "" {
		h.logger.Debugw("userIdentify", "error", errorEmptyHeader.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errorEmptyHeader.Error(),
		})

		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		h.logger.Debugw("userIdentify", "error", errorInvalidHeader.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errorInvalidHeader.Error(),
		})

		return
	}

	userID, err := h.userUsecase.ParseToken(headerParts[1])
	if err != nil {
		h.logger.Debugw("userIdentify", "error", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Set(userCtx, userID)
}
