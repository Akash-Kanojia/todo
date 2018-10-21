package main

import (
	"todo/internal/app/task"
	"todo/internal/app/user"
	"todo/internal/todo"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		task.Module,
		user.Module,
		todo.Module,
	)

	app.Run()
}
