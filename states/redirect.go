package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/find", redirect)
	http.HandleFunc("/root", root)
	http.HandleFunc("/barred", barred)
	http.HandleFunc("/redirect", redirect)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method inside redirect function", r.Method)
	w.Header().Set("Location", "/root")
	w.WriteHeader(http.StatusSeeOther)
	//w.WriteHeader(http.StatusTemporaryRedirect)
	//http.Redirect(w, r, "/barred", http.StatusMovedPermanently)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method inside root function", r.Method)
	fmt.Println(w)
	io.WriteString(w, "Page Redirected Here")
}

func barred(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)
}
