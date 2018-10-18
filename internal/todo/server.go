package todo

import (
	"fmt"
	"net/http"
	"todo/internal/app/task"
	"todo/internal/app/user"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Server struct {
	task task.Server
	user user.Server
}

func NewServer(
	task task.Server,
	user user.Server,
) Server {
	return Server{
		task: task,
		user: user,
	}
}

// NewMuxRouter is the convenience method to provide mux router.
func NewMuxRouter() *mux.Router {
	return mux.NewRouter()
}

// HttpServer listen and http request for todo apis.
func HttpServer() (err error) {

	go http.ListenAndServe(":80", context.ClearHandler(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s",
				r.Host,
				r.URL.String()),
				http.StatusMovedPermanently,
			)
		}),
	))

	return
}
