package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/puspa222/gopher-alert/internal/service"
)

type Handler struct {
	service *service.NotificationService
}

func NewHandler(s *service.NotificationService) *Handler {
	return &Handler{service: s}
}

type SendRequest struct {
	Provider string `json:"provider" binding:"required"`
	To       string `json:"to" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

func (h *Handler) Send(c *gin.Context) {
	var req SendRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Send(req.Provider, req.To, req.Message)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "sent",
	})
}
