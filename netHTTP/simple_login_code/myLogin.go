package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/federer", fedex)
	http.HandleFunc("/", sayhelloName)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form

	fmt.Println("Form of r", r.Form.Get("url_long")) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Printf("type of form		 is -- %T\n", r.Form)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func fedex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Federer won first round in Halle today")
	} else {
		fmt.Fprintf(w, "Federer will win Wimbledon")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Handler")
	if r.Method == "GET" {
		//parse the template file
		tmpl, err := template.ParseFiles("login.gohtml")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println(r.Form.Get("Age"))
	}

}
