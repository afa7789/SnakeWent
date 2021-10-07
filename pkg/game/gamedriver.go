package game

// Create new GameState and Start Game
func NewGame(w, h int) {
	initTerm()
	FirstLogic()
	g := createGameState(h, w)
	g.Print(false)
	for i := true; i; {
		i = g.roundIteration()
	}
}
