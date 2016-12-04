package main

import (
	"fmt"
	"net/rpc"

	"github.com/alexandremuzio/network_game/api"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("Error connecting")
		// handle error
	}
	args := &api.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	fmt.Println("reply = ", reply)
}
