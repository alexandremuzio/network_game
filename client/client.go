package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic("Error connecting")
		// handle error
	}
	//conn.Write([]byte("Send message to server\n"))
	fmt.Fprintf(conn, "Send message to server")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Printf("status = %v\n", status)
}
