package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/internal/app/user"
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
	var (
		task Task
		err  error
	)

	ctx := user.MakeContext(r)

	if err = json.NewDecoder(r.Body).Decode(&task); err == nil {
		if task, err = s.service.Create(ctx, task); err == nil {
			fmt.Println("Writing response")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		} else {
			w.Write([]byte(fmt.Sprintf("error in creating task, %v", err)))
		}
	} else {
		w.Write([]byte(fmt.Sprintf("invalid: task can't be decoded, %v", err)))
	}

}

// Update a task.
func (s Server) Update(w http.ResponseWriter, r *http.Request) {

}

// Find a task for given id.
func (s Server) Find(w http.ResponseWriter, r *http.Request) {

}

// FindAll tasks for given user.
func (s Server) FindAll(w http.ResponseWriter, r *http.Request) {

}

// Delete tasks for given id.
func (s Server) Delete(w http.ResponseWriter, r *http.Request) {

}
