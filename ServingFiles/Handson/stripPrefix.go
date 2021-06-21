package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/Handson/pics/", http.StripPrefix("/Handson/", http.FileServer(http.Dir("../Handson"))))
	log.Fatalln(http.ListenAndServe(":8000", nil))
}
