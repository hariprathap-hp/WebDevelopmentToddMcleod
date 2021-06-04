package main

import (
	"fmt"
	"html/template"
	"os"
)

type Month []struct {
	Name    string
	NumDays int
}

type Year struct {
	Yr      int
	Current Month
}

func main() {

	tmpl, err := template.ParseFiles("indexing.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	//pass a slice to template and access it based on index
	//sl := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September",
	//"October", "November", "December"}
	months := Month{
		{Name: "January", NumDays: 31}, {Name: "February", NumDays: 28},
		{Name: "March", NumDays: 31}, {Name: "April", NumDays: 30},
		{Name: "May", NumDays: 31}, {Name: "June", NumDays: 30},
		{Name: "July", NumDays: 31}, {Name: "August", NumDays: 31},
		{Name: "September", NumDays: 30}, {Name: "October", NumDays: 31},
		{Name: "November", NumDays: 30}, {Name: "December", NumDays: 31},
	}

	yr := Year{
		Yr:      2021,
		Current: months,
	}
	e_err := tmpl.Execute(os.Stdout, yr)
	if e_err != nil {
		fmt.Println(e_err)
	}
}
