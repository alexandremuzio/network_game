package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	_, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Recebi a msg: ", string(buf))
	conn.Write([]byte("Recebi a msg, manolo!"))
}

func main() {

	fmt.Printf("Hello, world.\n")

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		// handle error
		panic("Error!")
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}
