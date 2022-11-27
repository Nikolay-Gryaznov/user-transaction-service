package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user-transaction-service/appError"
	"user-transaction-service/pkg/model"
)

func (h *Handler) createUser(c *gin.Context) {
	var u model.User

	if err := c.BindJSON(&u); err != nil {
		appError.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.UserCreator.CreateUser(u)
	if err != nil {
		appError.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type getUsersResponse struct {
	Data []model.User `json:"data"`
}

func (h *Handler) getUsers(c *gin.Context) {
	users, err := h.services.UserReader.GetUsers()
	if err != nil {
		appError.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getUsersResponse{Data: users})
}

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		appError.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.UserReader.GetUserById(id)
	if err != nil {
		appError.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
