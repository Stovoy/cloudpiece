package cloudpiece

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting webserver on :80")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello. I'm speechless.")
}
