package game

// fillBorders changes int from board state to empty board
func fillBorders(sample *[][]int, width, height int) {
	// fill borders ,
	// for  i := range *sample {
	// 	for j, val := range (*sample)[i] {
	// 		if j == 0 || j == width+1 || i == 0 || i == height + 1 {
	// 			(*sample)[i][j] =
	// 		}
	// 	}
	// }
	for i := 1; i <= height; i++ {
		(*sample)[i][0] = cornerVertical
		(*sample)[i][height+1] = cornerVertical
	}

	for i := 1; i <= width; i++ {
		(*sample)[0][i] = cornerHorizontal
		(*sample)[width+1][i] = cornerHorizontal
	}

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
func fillWithCharacterFromNode(sample *[][]int, i int, nl nodeList) {
	//genericFunction to be used on fillFood and fillSnake
}

// fillFood, Change int on board for foods and check if adds more.
func fillFood(sample *[][]int, nl nodeList) {
	// check round number to see if can add one more food

	// receive a node with foods
	// while node != nil check it's position to the food one
	// change node pointer to next one.
}

// fillSnake, Change int on board for snake int
func fillSnake(sample *[][]int, nl nodeList) {
	// while node != nil
	// receive node , check it's position and change the gameboard Character to the snake one.
	// change node pointer to next node
}

func fillAll(g gameState) {
	fillEmptyDots(g.board, g.height, g.width)
	fillFood(g.board, g.foodList)
	fillSnake(g.board, g.snakeList)
}
