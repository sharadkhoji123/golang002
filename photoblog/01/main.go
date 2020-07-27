package main

import (
	"html/template"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	First    string
	Last     string
	Password []byte
	Contact  int
	Role     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, session

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/success", success)
	http.Handle("/encpic/public/", http.StripPrefix("/encpic/public", http.FileServer(http.Dir("./encpic/public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		password := req.FormValue("password")
		telephone := req.FormValue("phone")
		role := req.FormValue("role")

		if _, ok := dbUsers[username]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tel, err := strconv.Atoi(telephone)
		check(w, err)
		user := user{username, firstname, lastname, bs, tel, role}
		dbUsers[username] = user
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/success", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		u, ok := dbUsers[username]
		if !ok {
			http.Error(w, "Username does not exist", http.StatusForbidden)
			return
		}
		//does the entered password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username and Password do not match", http.StatusForbidden)
			return
		}
		c := getCookie(w, req)
		http.SetCookie(w, c)
		dbSessions[c.Value] = username
		http.Redirect(w, req, "/success", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	check(w, err)
	//delete the cookie
	delete(dbSessions, c.Value)
	//remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	tpl.ExecuteTemplate(w, "logout.gohtml", nil)

}
