package game

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

func hasEaten( next_p , actual_p position , fs nodeList ) bool{
	// check if next position is one of the foods position
	// remove food Node
	return false
}

func increaseSnakeSize(){
	// if it is increase snake size 
	// and add a node to it's last node.
	// change last node to new last node.
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

func checkIfNextPositionIsOK() bool{
	return false
}

func roundEnding(){
	// increase score for one more round, or print game end state.
	// not sure what else
}

func (g gameState) roundIteration() bool{
	// receive position
	pos := receivePosition();
	// see if has eaten, grow snake if eaten
	if ( hasEaten( pos, g.snakeHead , g.food ) ){
		increaseSnakeSize()
	}
	// move snake
	moveSnake()

	r := false
	// check if it's a valid position.
	// not snake body, not wall
	if ( checkIfNextPositionIsOK() ){
	// if ok , go next round
		roundEnding()
		r = true
	}else{
		r = false
	}

	// if not ok, end game.
	// draw 
	fillAll()

	return r
}
