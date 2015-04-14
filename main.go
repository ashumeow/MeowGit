package main

// GO Packages
import (
	"net/http"
	"html/template"
	"io/ioutil"
	"strings"
	"fmt"
)

// $git clone https://github.com/go-mgo/mgo.git
// using v2 == stable version
import (
	"MeowGit/mgo"
	"MeowGit/mgo/bson"
)

// MeowGit Packages
import (
	"MeowGit/users"
	"MeowGit/cache"
)
/*
import (
	"MeowGit/commits"
	"MeowGit/code"
)
*/
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
		http.Redirect(w, r, "/base", http.StatusFound)
		return
	} else if err != nil { 
		http.Redirect(w, r, "/_commits", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("_status.html")
	t.Execute(w, p)
}

// Handles the users loggin and gives them a cookie for doing so
func loginHandler(w http.ResponseWriter, r *http.Request) {

	usr := new(user.User)
	usr.Email = r.FormValue("email")
	pass := r.FormValue("pwd")
	userProfile := user.FindUser(usr.Email)

	if len(pass) > 0 {
		//usr.Password = code.SHA(pass)
		ok := user.CheckCredentials(usr.Email, usr.Password)
		if ok {
			usr = userProfile
			user.CreateUserFile(usr.Email) // TODO: Createuserfile
			cookie := cookies.LoginCookie(usr.Email)
			http.SetCookie(w, &cookie)
			usr.SessionID = cookie.Value
			_ = user.UpdateUser(usr)
			http.Redirect(w, r, "/login-succeeded", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login-failed", http.StatusFound)
		}
	} else {
		viewHandler(w, r)
	}
}

// Registers the new user
func registerHandler(w http.ResponseWriter, r *http.Request) {

	usr := new(user.User)
	usr.Email = r.FormValue("email")
	pass := r.FormValue("pwd")

	if len(pass) > 0 {
		//usr.Password = code.SHA(pass)
		if user.DoesAccountExist(usr.Email) {
			http.Redirect(w, r, "/account-exists", http.StatusFound)
		} else {
			ok := user.CreateAccount(usr)
			if ok {
				http.Redirect(w, r, "/success", http.StatusFound)
			} else {
				http.Redirect(w, r, "/failed", http.StatusFound)
			}
		}
	} else {
		viewHandler(w, r)
	}
}

// Logs out the user, removes their cookie from the database
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := new(user.User)
	sessionID := cookie.Value
	session, err := mgo.Dial("127.0.0.1:8080/")
	if err != nil {
		return
	}

	c := session.DB("test").C("users")
	c.Find(bson.M{"sessionid": sessionID}).One(&result)
	//result.SessionID = result.Email + ":" + code.SHA(result.SessionID+strconv.Itoa(rand.Intn(100000000)))
	err = c.Update(bson.M{"email": result.Email}, result)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}

func main() {

	// Adding function handlers
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)

	// Event Listener
    http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
