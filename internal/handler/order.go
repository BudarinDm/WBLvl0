package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"wblvl0/internal/model"
)

func (h *Handler) getOrderByUID(c *gin.Context) {
	var order model.Order

	uid := c.Param("uid")

	order, err := h.cache.GetInHandler(uid, h.service)
	if err != nil {

	}
	if reflect.ValueOf(order).IsZero() {
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
