package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	id     int
	title  string
	author string
}

func main() {
	p1 := Post{1, "Harry Potter", "J.K Rowling"}
	p2 := Post{2, "Game of Thrones", "George R.R Martin"}
	p3 := Post{3, "Lord of the Rings", "JJ Tolkien"}
	p4 := Post{4, "Winds of Winter", "George R.R Martin"}
	p5 := Post{5, "Fantastic Beasts and Where to Find Them", "J.K Rowling"}

	posts := []Post{p1, p2, p3, p4, p5}
	//create a csv file
	file1, err := os.Create("file.csv")
	if err != nil {
		fmt.Println("File creation failed")
	}
	defer file1.Close()
	writer := csv.NewWriter(file1)
	for _, post := range posts {
		line := []string{strconv.Itoa(post.id), post.title, post.author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	file2, err := os.Open("file.csv")
	if err != nil {
		panic(err)
	}

	defer file2.Close()
	reader := csv.NewReader(file2)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var reads []Post
	for _, item := range records {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{id: int(id), title: item[1], author: item[2]}
		reads = append(reads, post)
	}
	for _, p := range reads {
		fmt.Println(p)
	}
}
