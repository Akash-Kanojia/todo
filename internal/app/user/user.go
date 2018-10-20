package user

import (
	"context"
	"fmt"
)

type User struct {
	Email     string `json:"email"`
	Name      string `json:"Name"`
	APISecret string `json:"api_secret"`
}

func FromContext(ctx context.Context) (user User, err error) {
	var ok bool
	if user, ok = ctx.Value(key).(User); !ok {
		err = fmt.Errorf("Context doesn't have user")
	}
	return
}
