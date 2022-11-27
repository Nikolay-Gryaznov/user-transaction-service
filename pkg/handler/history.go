package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-transaction-service/appError"
	"user-transaction-service/pkg/model"
)

type getHistoryResponse struct {
	Data []model.History `json:"data"`
}

func (h *Handler) getAllHistory(c *gin.Context) {
	histories, err := h.services.HistoryReader.GetAllHistory()

	if err != nil {
		appError.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getHistoryResponse{Data: histories})
}

func (h *Handler) getHistoryByUserId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)

	if err != nil {
		appError.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	histories, err := h.services.HistoryReader.GetHistoryByUserId(id)
	if err != nil {
		appError.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getHistoryResponse{Data: histories})
}
