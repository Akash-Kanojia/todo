package task

import (
	"context"
	"fmt"
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
	if usr, ok := s.user.Auth(ctx); ok {
		if task, err = NewTask(raw.Title, raw.Body, usr.Email); err == nil {
			err = s.repo.Save(task)
		}
	} else {
		err = user.ErrInvalidUser
	}

	return
}

// Update a task.
func (s Service) Update(ctx context.Context, updateTask Task) (task Task, err error) {

	if _, ok := s.user.Auth(ctx); ok {
		if updateTask.ID == "" {
			err = fmt.Errorf("task id cannot be empty")
			return
		}

		if task, err = s.repo.Find(updateTask.ID); err != nil {
			return
		}

		task.Body = updateTask.Body
		task.Title = updateTask.Title
		task.Starred = updateTask.Starred
		task.UpdatedAt = time.Now()
		err = s.repo.Update(task)
	} else {
		err = user.ErrInvalidUser
	}

	return
}

// Find a task for given id.
func (s Service) Find(ctx context.Context, id string) (task Task, err error) {
	if usr, ok := s.user.Auth(ctx); ok {
		if task, err = s.repo.Find(id); err == nil {
			if task.UserEmail != usr.Email {
				err = user.ErrInvalidUser
			}
		}
	} else {
		err = user.ErrInvalidUser
	}
	return
}

// FindAll tasks for given user.
func (s Service) FindAll(ctx context.Context) (tasks []Task, err error) {
	if usr, ok := s.user.Auth(ctx); ok {
		tasks, err = s.repo.FindAll(usr)
	} else {
		err = user.ErrInvalidUser
	}
	return
}

// Delete tasks for given user.
func (s Service) Delete(ctx context.Context, ID string) (err error) {
	if _, ok := s.user.Auth(ctx); ok {
		if _, err = s.Find(ctx, ID); err == nil {
			err = s.repo.Delete(ID)
		}
	} else {
		err = user.ErrInvalidUser
	}
	return
}
