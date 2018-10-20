package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	emailKey  = "email"
	secretKey = "secret"
)

type Server struct {
	service Service
}

func NewServer(service Service) Server {
	return Server{
		service: service,
	}
}

// MakeContext retrieves email and secret from request header and creates context with user.
func MakeContext(r *http.Request) (ctx context.Context) {
	email := r.Header.Get(emailKey)
	secret := r.Header.Get(secretKey)

	user := User{
		Email:     email,
		APISecret: secret,
	}
	ctx = context.WithValue(r.Context(), key, user)

	return
}

// Create a user.
func (s Server) Create(w http.ResponseWriter, r *http.Request) {
	var (
		user User
		err  error
	)
	if err = json.NewDecoder(r.Body).Decode(&user); err == nil {
		if user, err = s.service.Create(user); err == nil {
			json.NewEncoder(w).Encode(user)
		} else {
			w.Write([]byte(fmt.Sprintf("error in creating user, %v", err)))
		}
	} else {
		w.Write([]byte("bad json: user can't be decoded."))
	}
}
