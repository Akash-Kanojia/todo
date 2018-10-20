package user

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collectionName = "user"

// Repository is an interface between persistence layer and user entity.
type Repository interface {
	Save(user User) error
	Find(email string) (User, error)
}

type RepositoryImpl struct {
	DB *mgo.Database
}

func NewRepository(db *mgo.Database) (impl Repository) {
	impl = &RepositoryImpl{
		DB: db,
	}
	return
}

// Save or update user.
func (repo *RepositoryImpl) Save(usr User) (err error) {
	err = repo.DB.C(collectionName).Insert(usr)
	return
}

// Find user by ID.
func (repo *RepositoryImpl) Find(email string) (usr User, err error) {
	err = repo.DB.C(collectionName).Find(bson.M{"email": email}).One(&usr)
	return
}
