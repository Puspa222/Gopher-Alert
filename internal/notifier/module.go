package notifier

import "go.uber.org/fx"

var Module = fx.Module("notifier",
	fx.Provide(NewRegistry,
		NewConsoleNotifier,
	))
