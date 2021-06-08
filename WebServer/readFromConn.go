package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	lis, lis_err := net.Listen("tcp", ":8000")

	if lis_err != nil {
		log.Fatalln(lis_err)
	}

	for {
		conn, con_err := lis.Accept()
		if con_err != nil {
			log.Panicln(con_err)
		}

		go connHandle(conn)
	}
}

func connHandle(c net.Conn) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	defer c.Close()
}
