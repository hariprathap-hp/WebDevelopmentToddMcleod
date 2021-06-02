package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Sage struct {
	Name  string
	Motto string
}

type Car struct {
	Brand   string
	Country string
}

type oxymoron struct {
	Sages []Sage
	Cars  []Car
}

var tpl *template.Template

//declate a funcmap
var ft = template.FuncMap{
	"uc": strings.ToUpper,
	"lt": lastThree,
}

var oxym oxymoron

func init() {
	tpl = template.Must(template.New("").Funcs(ft).ParseFiles("basicfunctemp.gohtml"))
}

func main() {
	sage1 := Sage{
		Name:  "Buddha",
		Motto: "Desire is the root cause of all miseries",
	}

	sage2 := Sage{
		Name:  "Matshona Dhliwayo",
		Motto: "Anger has great strength, but no brains",
	}

	car1 := Car{
		Brand:   "Audi",
		Country: "Germany",
	}

	car2 := Car{
		Brand:   "Toyota",
		Country: "Japan",
	}

	sages := []Sage{sage1, sage2}
	cars := []Car{car1, car2}

	oxym = oxymoron{sages, cars}

	http.HandleFunc("/functemp", funcHandler)
	log.Fatalln(http.ListenAndServe(":8000", nil))

}

func funcHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "basicfunctemp.gohtml", oxym)
	if err != nil {
		fmt.Println(err)
	}
}

func lastThree(s string) string {
	l := len(s)
	return s[l-3:]
}
