package handler

import (
	"net/http"
	"strconv"

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

	c.JSON(http.StatusOK, gin.H{
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
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	idList, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.Todolist.GetListById(userId, idList)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	idList, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todorest.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Todolist.UpdateList(userId, idList, input); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{
		Status: "updated",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	idList, err := strconv.Atoi(c.Param("id")) 
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Todolist.DeleteList(userId, idList); err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponce{Status: "success delete"})
}