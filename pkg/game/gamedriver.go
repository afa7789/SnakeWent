package game

import "github.com/afa7789/SnakeWent/pkg/printer"

// Create new GameState and Start Game
func NewGame(w, h int) {
	g := createGameState(h, w)
	hello := 0
	for i := true; i; {
		// reset()
		if hello == 0 {
			FirstLogic()
			hello = 1
		}
		// g.snakeList.Print()
		g.Print(true)
		i = g.roundIteration()
		if i {
			printer.PrintString("aqui ainda é true o retornado\n")
		} else {
			printer.PrintString("aqui não\n")
		}
	}
	// g.snakeList.Print()
	// closeTerm()
}
