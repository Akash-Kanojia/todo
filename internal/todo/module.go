package todo

import (
	"todo/internal/mongodb"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		mongodb.New,
	),

	fx.Invoke(
		HttpServer,
	),
)
