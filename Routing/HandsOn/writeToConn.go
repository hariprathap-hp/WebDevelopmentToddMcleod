package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	list, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		fmt.Println(err)
	}

	defer list.Close()

	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go serve(conn)
	}
}

func serve(c net.Conn) {
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		fmt.Println(text)
	}
	fmt.Println("code got here")
	io.WriteString(c, "This is my response to your request")
}
