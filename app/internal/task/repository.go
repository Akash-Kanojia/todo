package task

import (
	"todo/app/api/user"

	mgo "gopkg.in/mgo.v2"
)

// Repository is an interface between persistence layer and task entity.
type Repository interface {
	Create(task Task) error
	Update(task Task) error
	Find(ID string) Task
	FindAll(user user.User) []Task
	Delete(ID string) error
}

type MongoRepository struct {
	DB *mgo.Database
}

func NewRepository(db *mgo.Database) (impl *MongoRepository) {
	impl = &MongoRepository{
		DB: db,
	}
	return
}

func (repo *MongoRepository) Create(task Task) (err error) {

	return
}

func (repo *MongoRepository) Update(task Task) (err error) {

	return
}

func (repo *MongoRepository) Delete(ID string) (err error) {

	return
}

func (repo *MongoRepository) Find(ID string) (task Task) {

	return
}

func (repo *MongoRepository) FindAll(usr user.User) (tasks []Task) {

	return
}
