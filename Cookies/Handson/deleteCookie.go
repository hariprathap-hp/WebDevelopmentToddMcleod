package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set</a></h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	//create and set a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "del_cookie",
		Value: "delete",
	})

	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	//read a cookie. read a cookie and store it in a variable
	c, err := r.Cookie("del_cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<h1> Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("del_cookie")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	//Write the Max age value to the cookie in negative value
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
