package main

// GO Packages
import (
	"net/http"
	"html/template"
	"io/ioutil"
	"strings"
)

// MeowGit Packages
/* Also, includes
// $git clone https://github.com/ashumeow/mgo.git
// using v2.1 == stable version
*/
import (
	"MeowGit/users"
	"MeowGit/cache"
)

type Page struct {
	Title    string
	Body     template.HTML
	UserData template.HTML
	Bar      template.HTML
}

// Page Loading
func loadPage(title string, r *http.Request) (*Page, error) {

	filename, option, usr, bar := user.LoadUserInfo(title, r)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: template.HTML(body), UserData: (template.HTML(usr) + template.HTML(option)), Bar: template.HTML(bar)}, nil
}

// Showing a specific page
func viewHandler(w http.ResponseWriter, r *http.Request) {
	
	title := r.URL.Path[len("/"):]
	p, err := loadPage(title, r)

	z := strings.Split(title, "/")
	if z[0] == "commits" {
		http.ServeFile(w, r, title)
		return
	}

	if err != nil && !cookies.IsLoggedIn(r) {
		http.Redirect(w, r, "static/templates/base.html", http.StatusFound)
		return
	} else if err != nil { 
		http.Redirect(w, r, "static/templates/_commits.html", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("static/templates/_status.html")
	t.Execute(w, p)
}

func main() {

	// Adding function handlers
	http.HandleFunc("/", viewHandler)

	// Event Listener
    http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
