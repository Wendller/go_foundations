package main

import (
	"net/http"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", HomeHandler)
	// mux.Handle("/users", User{Name: "Wend"})
	fileServer := http.FileServer(http.Dir("./public"))
	fileServerMux := http.NewServeMux()

	fileServerMux.Handle("/", fileServer)

	// http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":8080", fileServerMux)
}

// Approach 1
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Mux"))
}

// Approach 2
type User struct {
	Name string
}

func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User: " + u.Name))
}
