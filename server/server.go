package server

import (
	"fmt"
	"net/http"
)

var staticDir string = "static/"

func Serve() {
	fmt.Println("Starting webserver on :80")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, staticDir+"index.html")
}
