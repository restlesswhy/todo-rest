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

	id, err := h.services.Item.CreateItem(userId, listId, itemInput)
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

	itemInput, err := h.services.Item.GetAllItems(userId, listId)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: itemInput,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input todorest.Item
	input, err = h.services.Item.GetItemById(userId, itemId)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var input todorest.UpdateItemInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Item.UpdateItem(userId, itemId, input); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	if err := h.services.Item.DeleteItem(userId, itemId); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{
		Status: "ok",
	})
}