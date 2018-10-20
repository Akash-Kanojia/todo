package task

import (
	"todo/internal/app/user"

	mgo "gopkg.in/mgo.v2"
)

// Repository is an interface between persistence layer and task entity.
type Repository interface {
	Save(task Task) error
	Find(ID string) (Task, error)
	FindAll(user user.User) ([]Task, error)
	Delete(ID string) error
}

// RepositoryImpl is the implementation of Repository.
type RepositoryImpl struct {
	DB *mgo.Database
}

// NewRepository returns an instance of RepositoryImpl.
func NewRepository(db *mgo.Database) (impl Repository) {
	impl = &RepositoryImpl{
		DB: db,
	}
	return
}

func (repo *RepositoryImpl) Find(ID string) (task Task, err error) {

	return
}

func (repo *RepositoryImpl) FindAll(usr user.User) (tasks []Task, err error) {

	return
}

// Save or update a task.
func (repo *RepositoryImpl) Save(task Task) (err error) {
	err = repo.DB.C("task").Insert(task)
	return
}

// Delete a task corresponding to given ID.
func (repo *RepositoryImpl) Delete(ID string) (err error) {

	return
}
