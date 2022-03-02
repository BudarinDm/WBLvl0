package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wblvl0/internal/model"
)

func (h *Handler) getOrderByUID(c *gin.Context) {
	var order model.Order

	uid := c.Param("uid")

	order, err := h.cache.GetInHandler(uid, h.service)
	if err != nil {

	}

	c.HTML(http.StatusOK, "getorder.tmpl", gin.H{
		"values": order,
	})
}
