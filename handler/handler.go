package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/restlesswhy/todo-rest/service"
)


type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

	api := r.Group("/api")
		{
			lists := api.Group("/lists", h.UserIdentity)
				{
					lists.POST("/", h.createList)
					lists.GET("/", h.getAllLists)
					lists.GET("/:id", h.getListById)
					lists.PUT("/:id", h.updateList)
					lists.DELETE("/:id", h.deleteList)

					item := lists.Group("/item")
						{
							item.POST("/", h.createItem)
							item.GET("/", h.getAllItems)
							item.GET("/:id", h.getItemById)
							item.PUT("/:id", h.updateItem)
							item.DELETE("/:id", h.deleteItem) 
						}
				}
		}
	return r
}