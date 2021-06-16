package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("requestValExpl.gohtml"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	tmpl.Execute(w, w.Header())
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("requestValExpl.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	p_err := r.ParseForm()
	if p_err != nil {
		fmt.Println(p_err)
	}

	myData := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		Contentlength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	w.WriteHeader(404)
	tmpl.Execute(w, myData)
}

func main() {
	http.HandleFunc("/formSubmit", formHandler)
	http.HandleFunc("/request", reqHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
