package todo

import (
	"todo/internal/app/task"
	"todo/internal/app/user"

	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		task.Module,
		user.Module,
	)

	app.Run()
}
