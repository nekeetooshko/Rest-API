package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader string = "Authorization"
)

// Идентификация борна (пользователя, вообще-то)
func (h *Handler) userIdentity(c *gin.Context) {

	// Дергаем значение из заголовка авторизации
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	// Хз пока зачем это
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "Incorrect auth header")
		return
	}

	// TODO: parse token
}
