package main

import (
	"fmt"
	"html/template"
	"math"
	"os"
	"time"
)

//declare two functions Sum and Sqrt
func Sum(a []int) float64 {
	s := 0
	for _, n := range a {
		s += n
	}
	return float64(s)
}

func Sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func formatDate(t time.Time) string {
	return t.Format(time.Kitchen)
}

//create a funcmap
var fm = template.FuncMap{
	"sum":     Sum,
	"sqrt":    Sqrt,
	"fmtDate": formatDate,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("pipelines.gohtml"))
}

func main() {
	val := []int{1, 2, 3, 4, 5}
	err := tpl.ExecuteTemplate(os.Stdout, "pipelines.gohtml", val)
	if err != nil {
		fmt.Println(err)
	}
}
