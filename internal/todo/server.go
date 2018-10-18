package todo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

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
