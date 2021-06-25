package main

import (
	"encoding/base64"
	"html/template"
	"io"
	"log"
	"net/http"
)

var dbUsers = map[string]string{}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	/*s := "What’s in a name? A rose by any other name would smell as sweet."
	be := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(be)
	bd, _ := base64.StdEncoding.DecodeString(be)
	fmt.Println(string(bd))

	fmt.Println()
	fmt.Println()
	fmt.Println("URL Encoding")

	ue := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(ue)*/

	http.HandleFunc("/", encode)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func encode(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		uname := r.FormValue("uname")
		pwd := r.FormValue("password")

		//first decode the password and compare it with entered password, if it matches, allow
		if ps, ok := dbUsers[uname]; ok {
			stored_pwd, _ := base64.StdEncoding.DecodeString(ps)
			if string(stored_pwd) == pwd {
				io.WriteString(w, "Welcome to the Shakespearean club!!!")
			} else {
				io.WriteString(w, "User Name or Password is Wrong")
			}
		} else {
			enc_pwd := base64.StdEncoding.EncodeToString([]byte(pwd))
			dbUsers[uname] = enc_pwd
		}
	}

	tmpl.ExecuteTemplate(w, "b64encode.gohtml", "Register to become a shakeapearean")
}
