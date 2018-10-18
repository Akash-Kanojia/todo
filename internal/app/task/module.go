package task

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewRepository,
		NewService,
		NewServer,
	),
)
