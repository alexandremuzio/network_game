package game

import "fmt"

func DisplayGameStart() {
	fmt.Println("################################")
	fmt.Println("Rock-Paper-Scissors")
	fmt.Println("################################")
	fmt.Println("Controls: ")
	fmt.Println("1 - Rock")
	fmt.Println("2 - Paper")
	fmt.Println("3 - Scissors")
	fmt.Println("")
	fmt.Println("Enjoy!! 😁")
}

func DisplayGameEnd(winnerId int) {
	fmt.Println("Game Ended!")
	fmt.Println("Player ", winnerId, " won!!")
	fmt.Println("################################")
}
