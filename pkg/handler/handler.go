package handler

import (
	"MaksJash/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service // Как я понял - это для доступа к БД
}

// Конструктор
func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

// Инициализация роутов
func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New() // Без ПО

	// Группа роутинга для авторизации/аутентификации
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// Группа роутинга для всего (и списков и элементов)
	api := router.Group("/api", h.userIdentity)
	{

		// Группа роутинга списков
		lists := api.Group("/lists") // userIdentity - это ПОшка для всех роутов
		{
			lists.GET("", h.getAlLists)
			lists.POST("", h.createList)
			lists.GET("/:id", h.getCertainList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// Группа роутинга элементов списков
			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAlItems)
				items.POST("/", h.createItem)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getCertainItem)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
