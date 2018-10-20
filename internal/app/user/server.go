package user

import (
	"context"
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

}
