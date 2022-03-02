package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getOrderByUID(c *gin.Context) {
	order, err := h.service.Order.GetOrder(c.Param("uid"))
	if err != nil {

	}

	c.HTML(http.StatusOK, "getorder.tmpl", gin.H{
		"values": order,
	})
}
