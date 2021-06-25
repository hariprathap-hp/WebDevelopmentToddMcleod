package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range gen() {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	time.Sleep(60 * time.Second)
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
		}
	}()
	return ch
}
