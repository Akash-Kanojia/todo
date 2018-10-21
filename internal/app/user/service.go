package user

import (
	"context"
	"fmt"

	"gopkg.in/mgo.v2"
)

const key = "user"

var ErrInvalidUser = fmt.Errorf("Invalid user credentials")

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

// Create a user.
func (s Service) Create(usr User) (user User, err error) {
	if _, err = s.repo.Find(usr.Email); err == mgo.ErrNotFound {
		user = NewUser(usr.Email, usr.Name)
		err = s.repo.Save(user)
	} else {
		err = fmt.Errorf("User already exist")
	}

	return
}

func (s Service) Auth(ctx context.Context) (user User, authentic bool) {
	var (
		ctxUser User
		err     error
	)

	if ctxUser, err = FromContext(ctx); err != nil {
		return
	}

	if len(ctxUser.Email)*len(ctxUser.APISecret) == 0 {
		return
	}

	if user, err = s.repo.Find(ctxUser.Email); err == nil {
		authentic = (user.APISecret == ctxUser.APISecret)
		return
	}

	return
}
