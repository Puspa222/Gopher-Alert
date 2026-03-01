package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"http",
	fx.Provide(NewHTTPServer, NewHandler),
	fx.Invoke(func(*gin.Engine) {}),
)
