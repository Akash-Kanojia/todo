package task

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes for task.
const (
	list   = "/list"
	task   = "/list/task"
	taskID = "/list/task/{id}"
)

func ServeRoutes(r *mux.Router, server Server) {
	r.HandleFunc(task, func(w http.ResponseWriter, r *http.Request) {
		server.Create(w, r)
	}).Methods("POST")

	r.HandleFunc(task, func(w http.ResponseWriter, r *http.Request) {
		server.Update(w, r)
	}).Methods("PUT")

	r.HandleFunc(taskID, func(w http.ResponseWriter, r *http.Request) {
		server.Find(w, r)
	}).Methods("GET")

	r.HandleFunc(list, func(w http.ResponseWriter, r *http.Request) {
		server.FindAll(w, r)
	}).Methods("GET")

	r.HandleFunc(taskID, func(w http.ResponseWriter, r *http.Request) {
		server.Delete(w, r)
	}).Methods("DELETE")

}
