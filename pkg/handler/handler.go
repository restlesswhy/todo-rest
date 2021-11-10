package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/restlesswhy/todo-rest/pkg/service"
)


type Handler struct {
	services service.Service 
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: *services}
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

					items := lists.Group(":id/items")
						{
							items.POST("/", h.createItem)
							items.GET("/", h.getAllItems)
						}
			}
			items := api.Group("/items", h.UserIdentity)
				{
					items.GET("/:id", h.getItemById)
					items.PUT("/:id", h.updateItem)
					items.DELETE("/:id", h.deleteItem) 
				}
		}
	return r
}