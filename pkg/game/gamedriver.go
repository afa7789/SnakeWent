package game

// Create new GameState and Start Game
func NewGame(w,h int) {
    // initTerm()
	// g := createGameState()
	// g.Print()
	initTerm()
	FirstLogic()
	g := createGameState(h,w)
	fillBorders(g.board,g.width,g.height)
	g.Print(true)
}

