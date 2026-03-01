package http

import (
	"github.com/gin-gonic/gin"
	"github.com/puspa222/gopher-alert/internal/service"
)

type Handler struct {
	service *service.NotificationService
}

func NewHandler(s *service.NotificationService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Get(c *gin.Context) {

}
