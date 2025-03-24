package handler

import (
	todo "MaksJash"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Для удобства ответа
type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// Ручки для списков

func (h *Handler) getAlLists(c *gin.Context) {

	// Достаем id пользователя, чтобы по нему вытащить все списки
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// Передача даннных на service
	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})

}

func (h *Handler) getCertainList(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	// Дергаем idшник списка из строки запроса
	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect id params")
		return
	}

	// Передача даннных на service
	list, err := h.services.TodoList.GetListById(user_id, list_id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) createList(c *gin.Context) {

	// Достаем id пользователя
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	// Взаимодейтсвие со списком
	var input todo.TodoList

	// Парсим список
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Передача даннных на service
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *Handler) updateList(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect id params")
		return
	}

	// Парсим обновленный список
	var input todo.UpdateListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Дергаем основную логику
	if err = h.services.TodoList.Update(user_id, list_id, input); err != nil { // Возможно здесь ошибка
		// У него вместо h.services.TodoList.Update - это: h.services.Update
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok!",
	})

}

func (h *Handler) deleteList(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	// Дергаем idшник списка из строки запроса, который необходимо удалить
	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect id params")
		return
	}

	// Передача даннных на service
	err = h.services.TodoList.DeleteListById(user_id, list_id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
