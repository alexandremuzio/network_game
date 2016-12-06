package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/alexandremuzio/network_game/game"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

//Read one number
func getInput() int {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	//exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b [1]byte
	os.Stdin.Read(b[0:])
	//fmt.Println("I got the byte", b, "("+string(b)+")")
	return int(b[0] - 48)
}

//Waits for server
func waitForTurn(conn net.Conn) int {
	var b [1]byte

	conn.Read(b[0:])

	//check if game ended
	reply := int(b[0])
	fmt.Println(reply)
	if reply == 4 {
		conn.Read(b[0:])
		return int(b[0])
	}

	return -1
}

//Sends move to server
func sendInputToServer(input int, conn net.Conn) {
	var b [1]byte

	b[0] = byte(input + 48)
	conn.Write(b[0:])
}

func main() {
	service := ":1201"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	game.DisplayGameStart()

	//Receive player id
	var b [1]byte
	conn.Read(b[0:])
	fmt.Println("I am player", int(b[0]))

	// player := NewClientPlayer(conn)

	input := -1
	for {
		//wait for my turn
		fmt.Println("Waiting for turn...")
		if result := waitForTurn(conn); result != -1 {
			//game ended
			game.DisplayGameEnd(result)
			os.Exit(0)
		}

		fmt.Println("Choose input...")
		//get valid input for game
		//1 - Rock
		//2 - Paper
		//3 - Scissors
		for {
			input = getInput()
			if input >= 1 && input <= 3 {
				//fmt.Println("Input is: ", input)
				break
			}
		}
		switch input {
		case 1:
			fmt.Println("Input: Rock 💎")
		case 2:
			fmt.Println("Input: Paper ♼")
		case 3:
			fmt.Println("Input: Scissors ✄")
		}
		sendInputToServer(input, conn)
	}

	// fmt.Fprintf(conn, "Send message to server")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// fmt.Printf("status = %v\n", status)

	// _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	// checkError(err)

	// //result, err := readFully(conn)
	// result, err := ioutil.ReadAll(conn)
	// checkError(err)

	// fmt.Println(string(result))

	os.Exit(0)
}
