package main

import "net"

type ClientPlayer struct {
	input int
	conn  net.Conn
}

func NewClientPlayer(conn net.Conn) *ClientPlayer {
	return &ClientPlayer{
		input: -1,
		conn:  conn,
	}
}

func (p *ClientPlayer) DoMove(input int) {
	p.input = input
}
