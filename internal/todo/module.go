package todo

import (
	"todo/internal/app/task"
	"todo/internal/app/user"
	"todo/internal/mongodb"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		mongodb.New,
		NewMuxRouter,
		NewServer,
	),

	fx.Invoke(
		HttpServer,
		task.ServeRoutes,
		user.ServeRoutes,
	),
)
