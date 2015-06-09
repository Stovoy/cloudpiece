package server

import (
	"fmt"
	"net/http"
	"flag"

	"github.com/gorilla/mux"
)

var hostname = flag.String("hostname", "localhost", "The hostname to use for routing.")

func Serve() {
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	app := r.Host(*hostname).Subrouter()
	app.HandleFunc("/", App)

	blog := r.Host("blog." + *hostname).Subrouter()
	blog.HandleFunc("/", Blog)
	blog.HandleFunc("/posts", Posts)

	http.Handle("/", r)

	fmt.Println("Starting webserver on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func App(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/app.html")
}

func Blog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/blog.html")
}

func Posts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `[{"title":"test"},{"title":"test 2"}]`)
}
