package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wblvl0/internal/model"
)

func (h *Handler) getOrderByUID(c *gin.Context) {
	var order model.Order

	uid := c.Param("uid")

	order, err := h.service.GetOrder(uid)
	if err != nil {
		c.HTML(http.StatusBadRequest, "errorForm.html", gin.H{
			"UID": uid,
		})
		return
	}

	c.HTML(http.StatusOK, "orderForm.html", gin.H{
		"UID":         order.UID,
		"TrackNumber": order.TrackNumber,
		"Address":     order.Delivery.Address,
		"Amount":      order.Payment.Amount,
	})
}
