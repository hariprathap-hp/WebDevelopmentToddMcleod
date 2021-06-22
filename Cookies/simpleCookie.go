package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set", setCookie)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/abundance", abundance)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "First_Cookie",
		Value: "Cookie_Defaulttt",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN. CHECK YOUR BROWSER")
	fmt.Fprintln(w, "In your browser, go to : dev tools/ application/ cookies")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("First_Cookie")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, "Your First Cookie", c)

	g, g_err := r.Cookie("general")

	if g_err != nil {
		fmt.Println(err)
	}

	s, s_err := r.Cookie("specific")
	if s_err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, "Your General Cookie", g)
	fmt.Fprintln(w, "Your Specific Cookie", s)
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "general_cookie",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "specific_cookie",
	})

	fmt.Fprintln(w, "COOKIE WRITTEN. CHECK YOUR BROWSER")
	fmt.Fprintln(w, "In your browser, go to : dev tools/ application/ cookies")
}
