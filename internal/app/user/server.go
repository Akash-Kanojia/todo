package user

import (
	"net/http"
)

type Server struct {
	service Service
}

func NewServer(service Service) Server {
	return Server{
		service: service,
	}
}

// Create a user.
func (s Server) Create(w http.ResponseWriter, r *http.Request) {

}
