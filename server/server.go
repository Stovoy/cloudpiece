package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	fmt.Println("Starting webserver on :80")
	r := mux.NewRouter()
	s := r.Host("localhost").Subrouter()
	s.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	s = r.Host("blog.localhost").Subrouter()
	s.PathPrefix("/").Handler(http.FileServer(http.Dir("./blog/static/")))
	http.Handle("/", r)
	http.ListenAndServe(":80", nil)
}
