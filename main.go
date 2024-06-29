package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home page</h1>")
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Teams!!</h1>")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/teams", teamsHandler)

	fmt.Println("Starting the server on port 3000...")
	http.ListenAndServe(":3000", nil)
}
