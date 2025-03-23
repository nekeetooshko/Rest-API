package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader string = "Authorization"
	userCtx             string = "userId"
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

	// Парсим токен
	id, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Если все ок - пишем значение id в контекст, дабы иметь к нему доступ в следующих ручках
	c.Set(userCtx, id)

}

// Достаем id пользователя из контекста
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "User id not found")
		return 0, errors.New("User id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "Cant convert id to int")
		return 0, errors.New("Cant convert id to int")
	}

	return idInt, nil
}
