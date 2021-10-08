package game

import "github.com/afa7789/SnakeWent/pkg/printer"

// Create new GameState and Start Game
func NewGame(w, h int) {
	g := createGameState(h, w)
	for i := true; i; {
		reset()
		FirstLogic()
		g.Print(true)
		i, g = g.roundIteration()
		if i {
			printer.PrintString("aqui ainda é true o retornado\n")
		} else {
			reset()
			FirstLogic()
			g.Print(true)
			// printer.PrintString("Hit Wall or Self\n")
			printer.PrintString("Game Ended")
		}
	}

}
