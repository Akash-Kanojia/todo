package todo

import (
	"flag"
	"fmt"
	"net/http"
	"todo/internal/app/config"
	"todo/internal/app/task"
	"todo/internal/app/user"

	"github.com/gorilla/context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	sslKey  = "/usr/ssl/key"
	sslCert = "/usr/ssl/cert"
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
func HttpServer(r *mux.Router) (err error) {
	localMode := flag.Bool("local", false, "set to true for running app locally")
	flag.Parse()

	if *localMode {
		go http.ListenAndServe(
			fmt.Sprintf(":%v", config.GetDef("PORT", "8090")),
			context.ClearHandler(handlers.CORS(
				handlers.AllowedMethods([]string{"DELETE", "GET", "POST", "PUT"}),
				handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Origin", "Accept"}),
				handlers.AllowCredentials(),
			)(r)))
	} else {
		// Expose http and https ports.
		go http.ListenAndServe(
			":80",
			context.ClearHandler(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, fmt.Sprintf("https://%s%s",
						r.Host,
						r.URL.String()),
						http.StatusMovedPermanently,
					)
				}),
			))

		http.ListenAndServeTLS(":443", sslCert, sslKey, context.ClearHandler(handlers.CORS(
			// handlers.AllowedOrigins([]string{"http://todo.api.com"}),
			handlers.AllowedMethods([]string{"DELETE", "GET", "POST", "PUT"}),
			handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Origin", "Accept"}),
			handlers.AllowCredentials(),
		)(r)))
	}

	return
}
