package handler

// Тут ответные ошибки

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	ErrMsg string `json:"message"`
}

// Типа удобной структуры для ответа
type statusResponse struct {
	Status string `json:"status"`
}

// Вернет ошибку и остановит все последующие ручки
func newErrorResponce(c *gin.Context, statusCode int, message string) {

	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode,
		Error{ErrMsg: message})

}
