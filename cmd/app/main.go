package main

import (
	"github.com/puspa222/gopher-alert/internal/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		notifier.Module,
		service.Module,
		config.Module,
	)

	app.Run()
}
