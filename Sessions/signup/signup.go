package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//create a struct type user
type User struct {
	Uname    string
	Password []byte
	Fname    string
	Lname    string
	Role     string
}

//Create the maps for dbSessions and dbUsers
var dbUsers = map[string]User{}
var dbSessions = map[string]string{}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob(("templates/*.gohtml")))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["hfiery@gmail.com"] = User{"hfiery@gmail.com", bs, "Hari", "Password", "007"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	//index function for what
	u := getUser(w, r)
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		http.Error(w, "Only the user 007 is allowed inside the bar", http.StatusForbidden)
		return
	}

	tmpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside signup")
	if alreadyLoggedin(r) {
		fmt.Println("Already Logged in")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		u := r.FormValue("uname")
		p := r.FormValue("password")
		f := r.FormValue("fname")
		l := r.FormValue("lname")
		role := r.FormValue("role")

		//If username already taken, forbid it
		if _, ok := dbUsers[u]; ok {
			http.Error(w, "USERNAME ALREADY TAKEN", http.StatusForbidden)
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = u

		//generate the passwords
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internel Server Error : ", http.StatusInternalServerError)
			return
		}
		usr := User{
			u, bs, f, l, role,
		}
		dbUsers[u] = usr

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		uname := r.FormValue("uname")
		password := r.FormValue("password")

		u, ok := dbUsers[uname]

		//This one is to elimmminate if theusername doesn't matchs
		if !ok {
			http.Error(w, "Username and/or Password donot match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username and/or Password donot match", http.StatusForbidden)
			return
		}

		//create a new session for the use and attach it to cookie
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = uname
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	tmpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedin(r) {
		http.Redirect(w, r, "User Not Logged In", http.StatusSeeOther)
		return
	}

	cookie, _ := r.Cookie("session")
	delete(dbSessions, cookie.Value)
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
