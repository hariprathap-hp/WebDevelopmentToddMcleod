package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/cricinfo/", cricHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func cricHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("cric.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, nil)
}
