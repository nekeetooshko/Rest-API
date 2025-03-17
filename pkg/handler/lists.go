package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ручки для списков

func (h *Handler) getAlLists(c *gin.Context) {
	c.Writer.Status()
}
func (h *Handler) getCertainList(c *gin.Context) {
}
func (h *Handler) createList(c *gin.Context) {

	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})

}
func (h *Handler) updateList(c *gin.Context) {
}
func (h *Handler) deleteList(c *gin.Context) {
}
