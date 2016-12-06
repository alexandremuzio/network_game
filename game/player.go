package game

import (
	"fmt"
	"net"
	"sync"
)

type Player interface {
	DoMove(wg *sync.WaitGroup)
	GetLastInput() int
	Lose()
}

type NetworkPlayer struct {
	id           int
	isTurn       bool
	input        int
	conn         net.Conn
	stillPlaying bool
}

func NewNetworkPlayer(id int, conn net.Conn) *NetworkPlayer {
	return &NetworkPlayer{
		id:           id,
		isTurn:       false,
		input:        -1,
		conn:         conn,
		stillPlaying: true,
	}
}

func (p *NetworkPlayer) DoMove(wg *sync.WaitGroup) {
	defer wg.Done()
	if !p.stillPlaying {
		return
	}

	//request input from client
	var b [1]byte
	p.conn.Write(b[0:])
	fmt.Println("Request sent")

	p.conn.Read(b[0:])
	p.input = int(b[0] - 48)
	fmt.Println("Input received", p.input)
	if p.input == 208 {
		p.Lose()
	}
}

func (p *NetworkPlayer) GetLastInput() int {
	return p.input
}
func (p *NetworkPlayer) Lose() {
	p.stillPlaying = false
}
