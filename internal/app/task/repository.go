package task

import (
	"todo/internal/app/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collectionName = "tasks"

// Repository is an interface between persistence layer and task entity.
type Repository interface {
	Save(task Task) error
	Update(task Task) error
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

// Find retrieves task for given id from database
func (repo *RepositoryImpl) Find(ID string) (task Task, err error) {
	err = repo.DB.C(collectionName).Find(bson.M{"id": ID}).One(&task)
	return
}

// FindAll retrieves all task for given user email from database
func (repo *RepositoryImpl) FindAll(usr user.User) (tasks []Task, err error) {
	err = repo.DB.C(collectionName).Find(bson.M{"useremail": usr.Email}).All(&tasks)
	return
}

// Save a task.
func (repo *RepositoryImpl) Save(task Task) (err error) {
	err = repo.DB.C(collectionName).Insert(task)
	return
}

// Update a task.
func (repo *RepositoryImpl) Update(task Task) (err error) {
	err = repo.DB.C(collectionName).Update(bson.M{"id": task.ID}, &task)
	return
}

// Delete a task corresponding to given ID.
func (repo *RepositoryImpl) Delete(ID string) (err error) {
	err = repo.DB.C(collectionName).Remove(bson.M{"id": ID})
	return
}
