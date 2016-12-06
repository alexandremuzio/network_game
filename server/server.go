package main

import (
	"fmt"
	"log"
	"net"

	"github.com/alexandremuzio/network_game/game"
)

func handleClient(gameEngine *game.GameEngine, conn net.Conn) {
	// close connection on exit
	// defer conn.Close()

	// Add new player
	gameEngine.AddPlayer(conn)

	//var buf [512]byte

	// read upto 512 bytes
	//n, err := conn.Read(buf[0:])
	//if err != nil {
	//	return
	//}

	// fmt.Fprintln(conn, "Recebi a msg: ", string(buf[:n]))
	//fmt.Fprintln(conn, "Recebi a msg: ", string(buf[:n]))
	//conn.Write([]byte("Recebi a msg, manolo!"))

	// // write the n bytes read
	// _, err2 := conn.Write(buf[0:n])
	// if err2 != nil {
	//     return
	// }
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

//Constants
const MaxPlayers int = 3

func main() {
	fmt.Println("Starting Game Server...")

	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// conn, err := listener.Accept()
	// var b [1]byte
	// conn.Read(b[0:])
	// fmt.Println(b)

	// var b1 [1]byte
	// conn.Read(b1[0:])
	// fmt.Println(b1)
	// fmt.Println("Im done")

	//start the game
	gameEngine := game.NewGameEngine()

	// fmt.Println(gameEngine)

	//Wait for all players to connect
	for {
		fmt.Println(gameEngine.NumPlayers())
		if gameEngine.NumPlayers() == MaxPlayers {
			break
		}

		conn, err := listener.Accept()
		fmt.Println("New player connected")
		if err != nil {
			continue
		}
		//fmt.Println(conn)
		//time.Sleep(5 * time.Second)
		//fmt.Fprintln(conn, "Hello m8")

		// run as a goroutine
		handleClient(gameEngine, conn)
	}

	//Play ball!
	for !gameEngine.Done() {
		fmt.Println(gameEngine)
		gameEngine.Step()
	}

	gameEngine.EndGame()
	game.DisplayGameEnd(gameEngine.Winner())
}
