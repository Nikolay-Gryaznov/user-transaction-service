package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-transaction-service/appError"
)

type transaction struct {
	Id     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (h *Handler) makeTransaction(c *gin.Context) {
	var t transaction

	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	result, err := h.services.TransactionMaker.MakeTransaction(t.Id, t.Amount)
	if err != nil {
		appError.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, result)
}
