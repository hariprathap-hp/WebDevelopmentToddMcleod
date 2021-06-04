package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(os.Stdout, 42)
}
