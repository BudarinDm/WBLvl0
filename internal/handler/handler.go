package handler

import (
	"github.com/gin-gonic/gin"
	"wblvl0/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("internal/web/templates/*")

	api := router.Group("/api/v1")
	{
		order := api.Group("/order")
		{
			order.GET("/search", h.searchOrder)
			order.GET("/", h.getOrderByUID)
		}
	}
	return router
}
