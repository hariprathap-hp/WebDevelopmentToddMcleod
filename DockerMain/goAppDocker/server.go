package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to GO Docker!!!")
}
