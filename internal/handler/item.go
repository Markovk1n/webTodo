package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	c.JSON(http.StatusOK, "create item")
}

func (h *Handler) getAllItems(c *gin.Context) {
	c.JSON(http.StatusOK, "get all item")
}

func (h *Handler) getItem(c *gin.Context) {
	c.JSON(http.StatusOK, "get item")
}

func (h *Handler) updateItem(c *gin.Context) {
	c.JSON(http.StatusOK, "update item")
}

func (h *Handler) deleteItem(c *gin.Context) {
	c.JSON(http.StatusOK, "delete item")
}
