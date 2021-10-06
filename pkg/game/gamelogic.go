package game

// receivePosition, receives from keyboard key arrow and returns new pos for game state.
func receivePosition(actual_pos,p_node_after_head position) position{

	// ver direção que está , x1 cabeça, x2 nodo seguinte
	// baixo x1 - x2 = 0 , y1 - y2 = -1
	// cima x1 - x2 = 0, y1 - y2 = 1
	// esquerda x1 - x2 = 1 , y1 - y2 = 0
	// direita x1 - x2 = -1 , y1 - y2 = 0
	x_diff := actual_pos.X - p_node_after_head.X
	y_diff := actual_pos.Y - p_node_after_head.Y

	var reject_actual_dir int

	if x_diff == 1{
		reject_actual_dir = arrowRight
	}else if x_diff == -1{
		reject_actual_dir = arrowLeft
	}else if y_diff == 1{
		reject_actual_dir = arrowDown
	}else if y_diff == -1{
		reject_actual_dir = arrowUp
	}

	// doesn't allow going to previous position, and receive next one
	var received int = receiveInput(reject_actual_dir)

	// new position reproposition
	if received == arrowLeft {
		actual_pos.X += 1
	}else if received == arrowRight {
		actual_pos.X -= 1
	}else if received == arrowUp{
		actual_pos.Y += 1
	}else if received == arrowDown{
		actual_pos.Y -= 1
	}

	// returns new position with addon/subtraction
	return position{
		X: actual_pos.X ,
		Y: actual_pos.Y ,
	}
}

// moveSnake receives nodeHead and it's next position, than call the subsequent nodes 
// with the actual position of it's parent for change.
func moveSnake(n *node,p position) {
	// checks if node pointer is nill,, return empty if it is
	if n == nil {
		return
	}
	// call this function for the next node with actual position
	moveSnake(n.nextNode,n.pos)
	// change the actual position for next position.
	n.pos = p
}

// Has Eaten ,receives position and food nodeList
func hasEaten( p position , fs nodeList , score *int) bool{

	has_haten := false
	// nodes for iteration
	var n_iter *node = fs.firstNode
	var n_previous *node = nil

	for ;n_iter !=nil;n_iter = n_iter.nextNode {
		// check if next position is one of the foods position
		if n_iter.pos == p {
			has_haten = true
			break
		}
		n_previous = n_iter
	}

	//has eaten is true, than delete the node from the list
	if has_haten {
		// remove food Node
		n_previous.nextNode = n_iter.nextNode
		*score += 10
	}
	// i believe the garbage collector will collecter previous iter.

	return has_haten
}

// IncreaseSnakeSize
// if it is increase snake size 
// and add a node to it's last node.
// change last node to new last node.
func increaseSnakeSize(pInt * int, snakeList * nodeList){
	*pInt += 1;
	str := "Snake"+intToString(*pInt);
	n := createNode(str,0,0,nil)
	snakeList.lastNode.nextNode = n
	snakeList.lastNode = n
}

func addFood(foodList * nodeList){
	// add a food to the nodeList of foods.
	n := createNode("food",0,0,nil)
	foodList.lastNode.nextNode = n
	foodList.lastNode = n
}

func headIsReturning() bool{
	// check next node from head position , if it's the same from next, if it is can't allow it for next pos.
	return false
}

func checkSelfHit(actual_pos position, snakeList * nodeList) bool{
	// check if snaker hit itself, if it does it has to lose.
	var n_iter *node = snakeList.firstNode

	for ;n_iter !=nil;n_iter = n_iter.nextNode {
		// check if next position is one of the foods position
		if n_iter.pos == actual_pos {
			return true
		}
	}

	return false
}

// check if next position hit's border
func checkBorderHit(actual_pos position,width,height int) bool{
	return false
}

// check if next position is a hit
func checkIfNextPositionIsOK(actual_pos position, snakeList * nodeList, w,h int) bool{
	if checkBorderHit(actual_pos,w,h) || checkSelfHit(actual_pos, snakeList)  {
		return false
	}
	return true
}

// roundEnding, not sure what'else do
func roundEnding(round,score *int){
	// increase score for one more round
	*score += 1
	*round += 1
	// not sure what else
}

func (g gameState) roundIteration() bool{
	// receive position
	pos := receivePosition(g.snakeHead,g.snakeList.firstNode.nextNode.pos); // new position
	// see if has eaten, grow snake if eaten
	if hasEaten( pos, g.foodList , &g.score) {
		increaseSnakeSize(&g.snakeLength,&g.snakeList)
	}
	// move snake
	moveSnake(g.snakeList.firstNode,pos)
	g.snakeHead = pos

	r := false // returned bool

	// check if it's a valid position.
	// not snake body, not wall
	if checkIfNextPositionIsOK() {
		// if ok , go next round
		roundEnding(&g.round,&g.score)
		r = true
	}else{
		r = false
	}

	// if not ok, end game.
	// draw 
	fillAll()

	return r
}