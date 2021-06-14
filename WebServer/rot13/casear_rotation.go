package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	list, list_err := net.Listen("tcp", ":8000")
	if list_err != nil {
		fmt.Println(list_err)
	}

	defer list.Close()

	for {
		conn, conn_err := list.Accept()
		if conn_err != nil {
			fmt.Println(conn_err)
		}
		go handle(conn)
		fmt.Println("Back from Handler")
	}
}

func handle(c net.Conn) {
	defer c.Close()
	//now need a scanner to read info from io stream

	scanner := bufio.NewScanner(c)
	scanner.Split(bufio.ScanWords)
	
	for scanner.Scan() {
		text := scanner.Text()
		//convert the string input to lower case
		l_text := strings.ToLower(text)
		//convert the string to byte stream to iterate it byte by byte
		bs := []byte(l_text)
		r := rot13(bs)
		fmt.Println(string(r))
		fmt.Fprintf(c, "Input is %s, and the caesar rotation value of the same is %s\n", l_text, r)
	}

}

func rot13(bs []byte) []byte {
	for i, v := range bs {
		if v <= 109 {
			bs[i] = v + 13
		} else {
			bs[i] = v - 13
		}
	}
	return bs
}
