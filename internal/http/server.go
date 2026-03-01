package http

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, handler *Handler) *gin.Engine {
	r := gin.Default()
	r.Use(APIKeyMiddleware())
	r.POST("v1/send", handler.Send)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go r.Run(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return r
}
