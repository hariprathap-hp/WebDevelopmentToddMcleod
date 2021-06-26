package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the index handler")
	_, err := io.WriteString(w, "This is for AWS")
	if err != nil {
		log.Println(err)
	}
}
