package storage

import "go.uber.org/fx"

var Module = fx.Module(
	"storage",
	fx.Provide(
		func() (*DB, error) {
			return NewSQLiteDB("notifications.db")
		},
	),
)
