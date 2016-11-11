package game

type IGameEngine interface {
	Start()
	AddPlayer()
	Step()
}

type GameEngine struct {
}

func (ge *GameEngine) Start() {
}

func (ge *GameEngine) AddPlayer() {
}

func (ge *GameEngine) Step() {
}
