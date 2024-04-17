package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/webTodo/pkg/models"
)

func (h *Handler) singUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	c.JSON(http.StatusOK, "Sing-up")
}

func (h *Handler) singIn(c *gin.Context) {
	c.JSON(http.StatusOK, "Sing-in")
}
