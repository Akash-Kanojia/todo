package user

import mgo "gopkg.in/mgo.v2"

// Repository is an interface between persistence layer and user entity.
type Repository interface {
	Create(user User) error
	Find(email string) User
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

func (repo *MongoRepository) Create(usr User) (err error) {

	return
}

func (repo *MongoRepository) Find(ID string) (usr User) {

	return
}
