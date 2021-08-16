package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyJson struct {
	ID        int
	Name      string
	Loc       Location
	Companies []Work
}

type Location struct {
	City  string
	State string
	Code  int
}

type Work struct {
	Name    string
	YOE     int
	Package float32
}

func main() {
	examp := MyJson{
		ID:   1,
		Name: "Hariprathap",
		Loc: Location{
			City:  "Chennai",
			State: "Tamilnadu",
			Code:  600100,
		},
		Companies: []Work{
			{Name: "Aricent", YOE: 6, Package: 8.5},
			{Name: "Aricent2", YOE: 3, Package: 29},
			{Name: "Aricent3", YOE: 1, Package: 31},
		},
	}

	o, e := json.MarshalIndent(&examp, "", "\t\t")
	if e != nil {
		fmt.Println("JSON Marshal failed")
	}
	fmt.Println(string(o))
	method2(examp)
}

func method2(ex MyJson) {
	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&ex)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
	fmt.Println(ex)
}
