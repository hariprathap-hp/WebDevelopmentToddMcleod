package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"time"
)

type t_info struct {
	Date time.Time
	Open float64
}

func main() {
	tmpl, err := template.ParseFiles("trade.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Open("trade.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	lines, r_err := csv.NewReader(f).ReadAll()
	if r_err != nil {
		fmt.Println(r_err)
	}

	tradeInfo := make([]t_info, 0, len(lines))
	for _, line := range lines {
		date, _ := time.Parse("2006-01-02", line[0])
		open, _ := strconv.ParseFloat(line[1], 64)
		t1 := t_info{
			Date: date,
			Open: open,
		}
		//append is a huge operation.If you know the capacity of the info beforehand, decare it before
		tradeInfo = append(tradeInfo, t1)
	}

	t_err := tmpl.Execute(os.Stdout, tradeInfo)
	if t_err != nil {
		fmt.Println(t_err)
	}
}
