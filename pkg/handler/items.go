package handler

import (
	todo "MaksJash"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Ручки для предметов списков

func (h *Handler) createItem(c *gin.Context) {

	// Id польз-ля
	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	// Дергаем idшник списка из строки запроса
	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect list id params")
		return
	}

	// Парсим
	var item todo.TodoItem
	if err = c.BindJSON(&item); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(user_id, list_id, item)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})

}

func (h *Handler) getAlItems(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect list id params")
		return
	}

	items, err := h.services.TodoItem.GetAll(user_id, list_id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)

}

func (h *Handler) getCertainItem(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	item_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect list id params")
		return
	}

	item, err := h.services.TodoItem.GetItemById(user_id, item_id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

func (h *Handler) updateItem(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect id params")
		return
	}

	var input todo.UpdateItemInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	// Дергаем основную логику
	if err = h.services.TodoItem.Update(user_id, list_id, input); err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok!",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {

	user_id, err := getUserId(c)
	if err != nil {
		return
	}

	item_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "Incorrect list id params")
		return
	}

	err = h.services.TodoItem.Delete(user_id, item_id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Okeeeee",
	})

}
