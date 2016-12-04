package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"

	"github.com/alexandremuzio/network_game/api"
)

type Arith int

func (t *Arith) Multiply(args *api.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *api.Args, quo *api.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	buf := make([]byte, 1024)

// 	_, err := conn.Read(buf)

// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}

// 	fmt.Println("Recebi a msg: ", string(buf))
// 	conn.Write([]byte("Recebi a msg, manolo!"))
// }

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
