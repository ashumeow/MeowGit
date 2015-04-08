package main

import (
	"net/http"
)

func main() {

	// Adding function handlers
	// For Testing the Go Build
	/*
	http.HandleFunc("/static/templates/base.html", base)
	http.HandleFunc("/static/templates/_commits.html", commits)
	http.HandleFunc("/static/templates/_status.html", status)
	http.HandleFunc("/static/templates/_details.html", details)
	*/

	// Event Listener
    http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
