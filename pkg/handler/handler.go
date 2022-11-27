package handler

import (
	"github.com/gin-gonic/gin"
	"user-transaction-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", h.getUsers)
			users.POST("/create", h.createUser)
			users.GET("/:id", h.getUserById)
		}
		api.POST("/transaction", h.makeTransaction)
		history := api.Group("/history")
		{
			history.GET("/", h.getAllHistory)
			history.GET("/:id", h.getHistoryByUserId)
		}
	}
	return router
}
