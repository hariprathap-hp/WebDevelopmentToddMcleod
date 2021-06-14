package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler int

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() == "/" {
		fmt.Fprintf(w, "No important url //")
	} else if r.URL.String() == "/something" {
		fmt.Fprintf(w, "URL given is \"Something\" ")
	}
}

func main() {
	var mh myHandler
	http.HandleFunc("/", simpleHandle)
	log.Fatalln(http.ListenAndServe(":8000", mh))
}

func simpleHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
