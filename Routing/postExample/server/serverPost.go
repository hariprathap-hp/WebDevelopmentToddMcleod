package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/post", postHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp))
}
