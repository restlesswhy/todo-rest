package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/restlesswhy/todo-rest"
)

type getAllItemsResponse struct {
	Data []todorest.Item `json:"data"`
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var itemInput todorest.Item
	if err := c.BindJSON(&itemInput); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Itemlist.CreateItem(userId, listId, itemInput)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	itemInput, err := h.services.Itemlist.GetAllItems(userId, listId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: itemInput,
	})
}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItem(c *gin.Context) {

}

func (h *Handler) deleteItem(c *gin.Context) {

}