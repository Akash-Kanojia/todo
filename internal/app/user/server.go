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

// Create a task.
func (s Server) Create(w http.ResponseWriter, r *http.Request) {

}

// Find a task for given id.
func (s Server) Find(w http.ResponseWriter, r *http.Request) {

}
