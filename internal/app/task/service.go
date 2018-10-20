package task

import (
	"context"
	"time"
	"todo/internal/app/user"
)

type Service struct {
	repo Repository
	user user.Service
}

func NewService(repo Repository, user user.Service) Service {
	return Service{
		repo: repo,
		user: user,
	}
}

// Create a task.
func (s Service) Create(ctx context.Context, raw Task) (task Task, err error) {
	if s.user.Auth(ctx) {
		if task, err = NewTask(raw.Title, raw.Body, raw.UserEmail); err == nil {
			err = s.repo.Save(task)
		}
	} else {
		err = user.ErrInvalidUser
	}

	return
}

// Update a task.
func (s Service) Update(ctx context.Context, updateTask Task) (err error) {
	var (
		task Task
	)

	if s.user.Auth(ctx) {
		if task, err = s.repo.Find(updateTask.ID); err != nil {
			return
		}

		task.Body = updateTask.Body
		task.Title = updateTask.Title
		task.Starred = updateTask.Starred
		task.UpdatedAt = time.Now()
		err = s.repo.Save(task)
	} else {
		err = user.ErrInvalidUser
	}

	return
}

// Find a task for given id.
func (s Service) Find(ctx context.Context, id string) (task Task, err error) {
	if s.user.Auth(ctx) {
		task, err = s.repo.Find(id)
	} else {
		err = user.ErrInvalidUser
	}
	return
}

// FindAll tasks for given user.
func (s Service) FindAll(ctx context.Context, usr user.User) (tasks []Task, err error) {
	if s.user.Auth(ctx) {
		tasks, err = s.repo.FindAll(usr)
	} else {
		err = user.ErrInvalidUser
	}
	return
}
