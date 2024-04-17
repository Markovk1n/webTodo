package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context) {
	c.JSON(http.StatusOK, "Sing-up")
}

func (h *Handler) singIn(c *gin.Context) {
	c.JSON(http.StatusOK, "Sing-in")
}
