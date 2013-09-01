package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", CoreHandler)

	http.Handle("/", r)

	fmt.Println("Running on port 3000...")
	http.ListenAndServe(":3000", nil)
}

func CoreHandler(w http.ResponseWriter, r *http.Request) {
	get(w, r, "Hello")
}

func get(w http.ResponseWriter, r *http.Request, b string) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, b)
}
