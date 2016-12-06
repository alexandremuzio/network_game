package game

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

//Implementation of a Rock Paper Scissor Game
type GameEngine struct {
	players []*NetworkPlayer
	done    bool
	winner  int
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		players: make([]*NetworkPlayer, 0),
	}
}

func (ge *GameEngine) Start() {
}

func (ge *GameEngine) AddPlayer(conn net.Conn) {
	var (
		id int
	)
	id = len(ge.players) + 1

	newPlayer := NewNetworkPlayer(id, conn)

	ge.players = append(ge.players, newPlayer)

	//send id to player
	b := []byte{byte(id)}
	conn.Write(b)
}

//game step
func (ge *GameEngine) Step() {
	if ge.gameEnded() {
		return
	}

	var wg sync.WaitGroup
	//Request move from players
	wg.Add(len(ge.players))
	fmt.Println("Number of ge.players: ", len(ge.players))
	for i, _ := range ge.players {
		go ge.players[i].DoMove(&wg)
	}
	wg.Wait()
	fmt.Println("Done waiting for players...")
	// time.Sleep(10 * time.Second)

	//Check if game ended
	rock := 0
	paper := 0
	scissors := 0

	for _, player := range ge.players {
		switch player.GetLastInput() {
		case 1:
			rock++
		case 2:
			paper++
		case 3:
			scissors++
		}
	}

	winner := -1
	//nothing happens
	if rock > 0 && paper > 0 && scissors > 0 {
		return
	}

	if rock > 0 && scissors > 0 {
		winner = 1
	} else if rock > 0 && paper > 0 {
		winner = 2
	} else if scissors > 0 && paper > 0 {
		winner = 3
	}
	fmt.Println("winner = ", winner)

	//update state
	for _, player := range ge.players {
		if winner != -1 &&
			player.GetLastInput() != winner {
			player.Lose()
		}
	}
}

func (ge *GameEngine) Done() bool {
	return ge.done
}

func (ge *GameEngine) gameEnded() bool {
	//check if game ended
	playersActive := 0
	activePlayerIdx := -1
	for idx, player := range ge.players {
		if player.stillPlaying {
			playersActive++
			fmt.Println("player active = ", idx)
			activePlayerIdx = idx
		}
	}
	if playersActive == 1 {
		ge.done = true
		ge.winner = activePlayerIdx + 1
	}
	return ge.done
}

func (ge *GameEngine) NumPlayers() int {
	return len(ge.players)
}

func (ge *GameEngine) Winner() int {
	return ge.winner
}

func (ge *GameEngine) EndGame() {
	b := [2]byte{4, byte(ge.winner)}
	for _, np := range ge.players {
		np.conn.Write(b[0:])
	}
}

func (ge *GameEngine) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "Players: %v\n", len(ge.players))
	for _, np := range ge.players {
		//np, _ := p.(*NetworkPlayer)
		fmt.Fprintf(&buffer, "Player %v: %v, %v\n",
			np.id, np.input, np.stillPlaying)
	}

	return buffer.String()
}
