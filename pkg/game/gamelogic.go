package game

// fillBorders changes int from board state to empty board
func fillBorders(){
   // fill borders , 
}

// fillEmptyDots, changes int from board state to empty board
func fillEmptyDots(){
	// fill everything that's not border
}

// fillWithCharacterFromNode, reutilization of code
func fillWithCharacterFromNode(){
	//genericFunction to be used on fillFood and fillSnake
}

// fillFood, Change int on board for foods and check if adds more.
func fillFood(){
	// check round number to see if can add one more food

	// receive a node with foods
	// while node != nil check it's position to the food one
	// change node pointer to next one.
}

// fillSnake, Change int on board for snake int
func fillSnake(){
	// while node != nil
	// receive node , check it's position and change the gameboard Character to the snake one.
	// change node pointer to next node
}

// receivePosition, receives from keyboard key arrow and returns new pos for game state.
func receivePosition() position{
	return position{
		X: 0,
		Y: 0,
	}
}

// moveSnake receives nodeHead and it's next position, than call the subsequent nodes 
// with the actual position of it's parent for change.
func moveSnake(){
	// checks if node pointer is nill,, return empty if it is
	// call this function for the next node with actual position
	// change the actual position for next position.
}

func hasEaten() bool{
	// check if next position is one of the foods position, if it is increase snake size 
	// and add a node to it's last node.
	// change last node to new last node.
	return false
}

func addFood() bool{
	// add a food to the nodeList of foods.
	return false
}

func headIsReturning() bool{
	// check next node from head position , if it's the same from next, if it is can't allow it for next pos.
	return false
}

func checkSelfHit() bool{
	// check if snaker hit itself, if it does it has to lose.
	return false
}

// check if next position hit's border
func checkBorderHit(i ,j,width,height int) bool{
	return false
}

func roundEnding(){
	// increase score for one more round, or print game end state.
	// not sure what else
}
