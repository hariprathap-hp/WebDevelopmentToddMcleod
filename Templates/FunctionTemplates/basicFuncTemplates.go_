package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

//declare a function template which has two string "uc" which is mapped to library function toUpper
//and a string "ft" which extracts the first 3 characters from a string and returns it
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	//tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("basicfunctemp.gohtml"))
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("basicfunctemp.gohtml"))
}

type Sage struct {
	Name  string
	Motto string
}

type Vehicle struct {
	Brand   string
	Country string
}

type Random struct {
	Wisdom []Sage
	Gaadi  []Vehicle
}

var sages []Sage

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[0:3]
}

func main() {
	b := Sage{
		Name:  "Buddha",
		Motto: "Desire is the root cause of all suffering",
	}

	g := Sage{
		Name:  "Gandhi",
		Motto: "Be the Change",
	}

	bnd := Sage{
		Name:  "JamesBond",
		Motto: "Bond, James Bond",
	}

	audi := Vehicle{
		Brand:   "Audi",
		Country: "Germany",
	}

	tata := Vehicle{
		Brand:   "tata",
		Country: "India",
	}

	sages = []Sage{b, g, bnd}
	vehicles := []Vehicle{audi, tata}

	randoms := Random{sages, vehicles}

	err := tpl.ExecuteTemplate(os.Stdout, "basicfunctemp.gohtml", randoms)
	if err != nil {
		fmt.Println("Error is -- ", err)
	}
	//http.HandleFunc("/functemp/", funcHandler)
	//log.Fatalln(http.ListenAndServe(":8000", nil))
}

func funcHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, sages)
	if err != nil {
		fmt.Println("Error is ", err)
	}
}
