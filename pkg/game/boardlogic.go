package game

// fillBorders changes int from board state to empty board
func fillBorders(sample *[][]int, width, height int) {
	// vertical bars
	for i := 1; i <= height; i++ {
		(*sample)[i][0] = cornerVertical
		(*sample)[i][height+1] = cornerVertical
	}
	//horizontal bars
	for i := 1; i <= width; i++ {
		(*sample)[0][i] = cornerHorizontal
		(*sample)[width+1][i] = cornerHorizontal
	}
	//corners
	(*sample)[0][0] = cornerUpLeft
	(*sample)[0][width+1] = cornerUpRight
	(*sample)[height+1][0] = cornerDownLeft
	(*sample)[height+1][width+1] = cornerDownRight
}

// fillEmptyDots, changes int from board state to empty board
func fillEmptyDots(sample *[][]int, height, width int) {
	// fill everything that's not border
	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			(*sample)[i][j] = dot
		}
	}

}

// fillWithCharacterFromNode, reutilization of code
func fillWithCharacterFromNode(sample *[][]int, nl nodeList, i int) {
	//genericFunction to be used on fillFood and fillSnake
	// while node != nil
	// receive node , check it's position and change the gameboard Character to the snake one.
	// change node pointer to next node
}

// fillFood, Change int on board for foods and check if adds more.
func fillFood(sample *[][]int, nl nodeList) {
	// call the function and fills with diamondEmpty int
	fillWithCharacterFromNode(sample, nl, diamondEmpty)
}

// fillSnake, Change int on board for snake int
func fillSnake(sample *[][]int, nl nodeList) {
	// call the function and fills with rectangle int
	fillWithCharacterFromNode(sample, nl, rectangle)
}

// Call all fills, that will be called more than one time.
func fillAll(g gameState) {
	fillEmptyDots(g.board, g.height, g.width)
	fillFood(g.board, g.foodList)
	fillSnake(g.board, g.snakeList)
}
