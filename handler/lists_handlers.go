package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todorest "github.com/restlesswhy/todo-rest"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	var listInput todorest.List
	if err := c.BindJSON(&listInput); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Todolist.CreateList(userId, listInput)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListResponse struct {
	Data []todorest.List `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.services.Todolist.GetAllLists(userId)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}