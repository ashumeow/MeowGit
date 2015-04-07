package main

import (
	"net/http"
	"text/template"
)

type Template struct {
	// add something
}

func base(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		//t := template.New("base.html")
		t, r := template.ParseFiles("base.html")
		t.Execute(w, r)
	}
}

func commits(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		//commit := pushCommits()
		//t := template.New("_commits.html")
		t, r := template.ParseFiles("_commits.html")
		t.Execute(w, r)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		//t := template.New("_status.html")
		t, r := template.ParseFiles("_status.html")
		t.Execute(w, r)
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		//t := template.New("_details.html")
		t, r := template.ParseFiles("_details.html")
		t.Execute(w, r)
	}
}

func diff(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func pull(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func push(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func checkout(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func create_branch(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func diff_stat(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func stage(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func unstage(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func get_heads_with_commit(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func get_branches_in_remote(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func merge(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func open(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func run(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func stop(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func stopped(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func exit(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func only_server(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path[1:] == "") {
		// add something
	}
}

func terminate_gitj_sig() {
	//terminate_gitj()
}

func terminate_gitj(w http.ResponseWriter, r *http.Request) {
	//exit()
}

func main() {

	// Adding function handlers
	// For Testing the Go Build
	http.HandleFunc("/static/templates/base.html", base)
	http.HandleFunc("/static/templates/_commits.html", commits)
	http.HandleFunc("/static/templates/_status.html", status)
	http.HandleFunc("/static/templates/_details.html", details)


	// TODO: Add more...

	// Event Listener
    http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
