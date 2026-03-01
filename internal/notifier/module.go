package notifier

import "go.uber.org/fx"

var Module = fx.Module(
	"notifier",
	fx.Provide(
		fx.Annotate(
			NewConsoleNotifier,
			fx.As(new(Notifier)), // bind to interface
			fx.ResultTags(`group:"notifiers"`),
		),
		fx.Annotate(
			NewDiscordNotifier,
			fx.As(new(Notifier)),
			fx.ResultTags(`group:"notifiers"`),
		),
		fx.Annotate(
			NewRegistry,
			fx.ParamTags(`group:"notifiers"`),
		),
	),
)
