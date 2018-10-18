package task

import (
	"time"
	"todo/internal/app/user"
)

type Service struct {
	repo *RepositoryImpl
}

func NewService(repo *RepositoryImpl) Service {
	return Service{
		repo: repo,
	}
}

// Create a task.
func (s Service) Create(task Task) (err error) {
	err = s.repo.Save(task)
	return
}

// Update a task.
func (s Service) Update(updateTask Task) (err error) {
	var (
		task Task
	)

	if task, err = s.repo.Find(updateTask.ID); err != nil {
		return
	}

	task.Body = updateTask.Body
	task.Title = updateTask.Title
	task.Starred = updateTask.Starred
	task.UpdatedAt = time.Now()

	err = s.repo.Save(task)
	return
}

// Find a task for given id.
func (s Service) Find(id string) (task Task, err error) {
	task, err = s.repo.Find(id)
	return
}

// FindAll tasks for given user.
func (s Service) FindAll(usr user.User) (tasks []Task, err error) {
	tasks, err = s.repo.FindAll(usr)
	return
}
