package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Email     string `json:"email"`
	Name      string `json:"Name"`
	APISecret string `json:"api_secret"`
}

func NewUser(
	email string,
	name string,
) (user User) {
	user = User{
		Email:     email,
		Name:      name,
		APISecret: uuid.New().String(),
	}
	return
}

func FromContext(ctx context.Context) (user User, err error) {
	var ok bool
	if user, ok = ctx.Value(key).(User); !ok {
		err = fmt.Errorf("Context doesn't have user")
	}
	return
}
