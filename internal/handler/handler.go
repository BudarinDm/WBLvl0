package handler

import (
	"github.com/gin-gonic/gin"
	"wblvl0/internal/cache"
	"wblvl0/internal/service"
)

type Handler struct {
	service *service.Service
	cache   *cache.Cache
}

func NewHandler(service *service.Service, cache *cache.Cache) *Handler {
	return &Handler{service: service, cache: cache}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("internal/web/templates/*")

	api := router.Group("/api/v1")
	{
		order := api.Group("/order")
		{
			order.GET("/:uid", h.getOrderByUID)
		}
	}
	return router
}
