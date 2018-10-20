package user

import (
	"context"
	"fmt"
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
func (s Service) Create(user User) (err error) {
	err = s.repo.Save(user)
	return
}

// Find a user for given email.
func (s Service) Find(email string) (user User, err error) {
	user, err = s.repo.Find(email)
	return
}

func (s Service) Auth(ctx context.Context) (authentic bool) {
	var (
		ctxUser, usr User
		err          error
	)

	if ctxUser, err = FromContext(ctx); err != nil {
		return
	}

	if len(ctxUser.Email)*len(ctxUser.APISecret) == 0 {
		return
	}

	if usr, err = s.repo.Find(ctxUser.Email); err == nil {
		return usr.APISecret == ctxUser.APISecret
	}
	return
}
