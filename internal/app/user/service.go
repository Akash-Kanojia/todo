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
func (s Service) Create(usr User) (user User, err error) {
	user = NewUser(usr.Email, usr.Name)
	err = s.repo.Save(user)
	return
}

// Find a user for given email.
func (s Service) Find(email string) (user User, err error) {
	user, err = s.repo.Find(email)
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
