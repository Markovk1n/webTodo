package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	c.JSON(http.StatusOK, "create List")
}

func (h *Handler) getAllLists(c *gin.Context) {
	c.JSON(http.StatusOK, "get all list")
}

func (h *Handler) getList(c *gin.Context) {
	c.JSON(http.StatusOK, "get list")
}

func (h *Handler) updateList(c *gin.Context) {
	c.JSON(http.StatusOK, "update list")
}

func (h *Handler) deleteList(c *gin.Context) {
	c.JSON(http.StatusOK, "delete list")
}
