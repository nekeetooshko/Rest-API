package handler

import (
	todo "MaksJash"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Авторизационные/регистрационные ручки

// Регистрация
func (h *Handler) signUp(c *gin.Context) {
	// Для реги нужны name, username и password

	userData := todo.User{} // Сюда будут записаны данные с клиента

	if err := c.BindJSON(&userData); err != nil { // Парсит с ответом клиенту. Зачем-то

		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// После парсинга и валидации тела запроса - передача данных на слой ниже - в сервис, на котором реализована
	// логика регистрации
	id, err := h.services.Authorization.CreateUser(userData)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})

}

// Авторизация
func (h *Handler) signIn(c *gin.Context) {

}
