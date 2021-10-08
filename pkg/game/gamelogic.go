package game

import (
	"math/rand"
)

// receivePosition, receives from keyboard key arrow and returns new pos for game state.
func receivePosition(actual_pos, p_after position) position {

	// actual_pos.Print()
	// p_after.Print()
	// ver direção que está , x1 cabeça, x2 nodo seguinte
	// baixo x1 - x2 = 0 , y1 - y2 = -1
	// cima x1 - x2 = 0, y1 - y2 = 1
	// esquerda x1 - x2 = 1 , y1 - y2 = 0
	// direita x1 - x2 = -1 , y1 - y2 = 0
	x_diff := actual_pos.X - p_after.X
	y_diff := actual_pos.Y - p_after.Y
	// printer.PrintString("\n\t" + intToString(x_diff) + "\n")
	// printer.PrintString(intToString(y_diff) + "\n")

	var reject_actual_dir int

	if x_diff == 1 {
		// printer.PrintString("nao pode esquerda\n")
		reject_actual_dir = arrowLeft
	} else if x_diff == -1 {
		// printer.PrintString("nao pode direita\n")
		reject_actual_dir = arrowRight
	} else if y_diff == 1 {
		// printer.PrintString("nao pode cima\n")
		reject_actual_dir = arrowUp
	} else if y_diff == -1 {
		// printer.PrintString("nao pode baixo\n")
		reject_actual_dir = arrowDown
	}

	// doesn't allow going to previous position, and receive next one
	var received int = receiveInput(reject_actual_dir)

	// new position reproposition
	if received == arrowLeft {
		// printer.PrintString("arrow left\n")
		actual_pos.X -= 1
	} else if received == arrowRight {
		// printer.PrintString("arrow right\n")
		actual_pos.X += 1
	} else if received == arrowUp {
		// printer.PrintString("arrow up\n")
		actual_pos.Y -= 1
	} else if received == arrowDown {
		// printer.PrintString("arrow right\n")
		actual_pos.Y += 1
	}

	// returns new position with addon/subtraction
	return position{
		X: actual_pos.X,
		Y: actual_pos.Y,
	}
}

// moveSnake receives nodeHead and it's next position, than call the subsequent nodes
// with the actual position of it's parent for change.
func moveSnake(n *node, p position) {
	// checks if node pointer is nill,, return empty if it is
	if n == nil {
		return
	}
	// call this function for the next node with actual position
	moveSnake(n.nextNode, n.pos)
	// change the actual position for next position.
	n.pos = p
}

// Has Eaten ,receives position and food nodeList
func hasEaten(p position, fs *nodeList, score *int) bool {

	// nodes for iteration
	var n_iter *node = fs.firstNode
	var n_previous *node = nil

	for ; n_iter != nil; n_iter = n_iter.nextNode {
		// check if next position is one of the foods position
		if n_iter.pos == p {
			// remove food Node
			*score += 10
			// verifica se ja saimos do primeiro nodo
			if n_previous != nil {
				n_previous.nextNode = n_iter.nextNode
			} else {
				// é o primeiro nodo que foi comido so pula a lista para o próximo e ignora o primeiro.
				fs.firstNode = n_iter.nextNode
			}

			return true
		}
		n_previous = n_iter
	}

	// i believe the garbage collector will collecter previous iter.
	return false
}

// IncreaseSnakeSize
// if it is increase snake size
// and add a node to it's last node.
// change last node to new last node.
func increaseSnakeSize(pInt *int, snakeList *nodeList) {
	*pInt += 1
	str := "Snake" + intToString(*pInt)
	n := createNode(str, 0, 0, nil)
	snakeList.lastNode.nextNode = n
	snakeList.lastNode = n
}

// add a food to the nodeList of foods.
func addFood(foodList *nodeList, w, h int) {
	n := createNode("food", rand.Intn(w-1)+1, rand.Intn(h-1)+1, nil)

	if foodList.firstNode != nil {
		if foodList.lastNode == nil {
			foodList.firstNode.nextNode = n
			foodList.lastNode = n
		} else {
			foodList.lastNode.nextNode = n
		}
		foodList.lastNode = n
	} else {
		foodList.firstNode = n
	}

}

// func headIsReturning() bool { CANCELED , made in first function
// 	// check next node from head position , if it's the same from next, if it is can't allow it for next pos.
// 	return false
// }

func checkSelfHit(actual_pos position, snakeList *nodeList) bool {
	// check if snaker hit itself, if it does it has to lose.
	var n_iter *node = snakeList.firstNode.nextNode

	for ; n_iter != nil; n_iter = n_iter.nextNode {
		// check if next position is one of the foods position
		if n_iter.pos == actual_pos {
			return true
		}
	}

	return false
}

// check if next position hit's border MISSING
func checkBorderHit(actual_pos position, width, height int) bool {
	return false
}

// check if next position is a hit
func checkIfNextPositionIsOK(actual_pos position, snakeList *nodeList, w, h int) bool {

	if checkBorderHit(actual_pos, w, h) || checkSelfHit(actual_pos, snakeList) {
		return false
	}
	return true
}

// roundEnding, not sure what'else do
func roundEnding(round, score *int) bool {
	*round = *round + 1
	// increase score each 5 rounds
	if *round%5 == 0 {
		*score = *score + 1
	}
	// on random number
	var random1 int = rand.Intn(*score+1) + *round
	var random2 int = rand.Intn(int(*score/6)+1) + 5
	return random1%random2 == 0
}

func (g gameState) roundIteration() (bool, gameState) {
	// receive position
	pos := receivePosition(g.snakeList.firstNode.pos, g.snakeList.firstNode.nextNode.pos) // new position

	// see if has eaten, grow snake if eaten
	if hasEaten(pos, &g.foodList, g.score) {
		increaseSnakeSize(g.snakeLength, &g.snakeList)
	}
	// move snake
	moveSnake(g.snakeList.firstNode, pos)
	g.snakeHead = g.snakeList.firstNode.pos

	r := false //returned bool

	// check if it's a valid position.
	// not snake body, not wall
	if checkIfNextPositionIsOK(pos, &g.snakeList, g.width, g.height) {
		// if ok , go next round
		if roundEnding(g.round, g.score) {
			addFood(&g.foodList, g.width, g.height)
			// g.foodList.Print()
		}
		r = true
	} else {
		r = false
	}

	// if not ok, end game.
	// draw
	fillAll(g)

	return r, g
}
