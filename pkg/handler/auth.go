package handler

import (
	todo "MaksJash"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Авторизационные/регистрационные ручки

// @Summary Client sign up
// @Tags Authentification endpoints
// @Description create user account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-up [post]
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

// Структура для авторизации (обычный todo.User не подойдет)
type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Client sign in
// @Tags Authentification endpoints
// @Description login into user account
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) { // Авторизация
	// Для авторизации нужны username и password

	var userData signInInput

	if err := c.BindJSON(&userData); err != nil {

		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// После парсинга и валидации тела запроса - передача данных на слой ниже - в сервис, на котором реализована
	// логика регистрации
	token, err := h.services.Authorization.GenerateToken(userData.Username, userData.Password)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
