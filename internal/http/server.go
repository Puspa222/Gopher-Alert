package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle, handler *Handler) *gin.Engine {
	r := gin.Default()
	r.Use(APIKeyMiddleware())
	r.POST("/v1/send", handler.Send)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				fmt.Println("Starting server on:8080")
				if err := r.Run(":8080"); err != nil {
					fmt.Printf("Server error:%v", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Printf("\nStopping Server")
			return nil
		},
	})

	return r
}
