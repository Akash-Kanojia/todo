package user

import mgo "gopkg.in/mgo.v2"

// Repository is an interface between persistence layer and user entity.
type Repository interface {
	Save(user User) error
	Find(email string) (User, error)
}

type RepositoryImpl struct {
	DB *mgo.Database
}

func NewRepository(db *mgo.Database) (impl *RepositoryImpl) {
	impl = &RepositoryImpl{
		DB: db,
	}
	return
}

// Save or update user.
func (repo *RepositoryImpl) Save(usr User) (err error) {

	return
}

// Find user by ID.
func (repo *RepositoryImpl) Find(ID string) (usr User, err error) {

	return
}
