package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes for task.
const (
	signup = "user/signup"
)

func Listen(r *mux.Router, server Server) {
	r.HandleFunc(signup, func(w http.ResponseWriter, r *http.Request) {
		server.Create(w, r)
	}).Methods("POST")
}
