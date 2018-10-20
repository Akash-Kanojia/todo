package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/internal/app/user"

	"github.com/gorilla/mux"
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
		w.Write([]byte(fmt.Sprintf("bad json: task can't be decoded, %v", err)))
	}

}

// Update a task.
func (s Server) Update(w http.ResponseWriter, r *http.Request) {
	var (
		task Task
		err  error
	)

	ctx := user.MakeContext(r)

	if err = json.NewDecoder(r.Body).Decode(&task); err == nil {
		if task, err = s.service.Update(ctx, task); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		} else {
			w.Write([]byte(fmt.Sprintf("error in updating task, %v", err)))
		}
	} else {
		w.Write([]byte(fmt.Sprintf("bad json: task can't be decoded, %v", err)))
	}
}

// Find a task for given id.
func (s Server) Find(w http.ResponseWriter, r *http.Request) {
	var (
		task Task
		err  error
	)

	ctx := user.MakeContext(r)

	if id := mux.Vars(r)["id"]; id != "" {
		if task, err = s.service.Find(ctx, id); err == nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		} else {
			w.Write([]byte(fmt.Sprintf("error in fetching task for id %v, %v", id, err)))
		}
	} else {
		w.Write([]byte(fmt.Sprintf("bad query: id not found, %v", err)))
	}
}

// FindAll tasks for given user.
func (s Server) FindAll(w http.ResponseWriter, r *http.Request) {
	var (
		tasks []Task
		err   error
	)

	ctx := user.MakeContext(r)

	if tasks, err = s.service.FindAll(ctx); err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	} else {
		w.Write([]byte(fmt.Sprintf("error in fetching list of task, %v", err)))
	}

}

// Delete tasks for given id.
func (s Server) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	ctx := user.MakeContext(r)

	if id := mux.Vars(r)["id"]; id != "" {
		if err = s.service.Delete(ctx, id); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Successfully deleted task"))
		} else {
			w.Write([]byte(fmt.Sprintf("error in deleting task for id %v, %v", id, err)))
		}
	} else {
		w.Write([]byte(fmt.Sprintf("bad query: id not found, %v", err)))
	}
}
