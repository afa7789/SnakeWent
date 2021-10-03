package game

import(
	"os"
    "os/exec"
	"fmt"
)

// receivePosition, receives from keyboard key arrow and returns new pos for game state.
func receivePosition(actual_pos,p_node_after_head position) position{

	return position{
		X: 0,
		Y: 0,
	}
}

// moveSnake receives nodeHead and it's next position, than call the subsequent nodes 
// with the actual position of it's parent for change.
func moveSnake(n *node,p position) {
	// checks if node pointer is nill,, return empty if it is
	if (n == nil ){
		return
	}
	// call this function for the next node with actual position
	moveSnake(n.nextNode,n.pos)
	// change the actual position for next position.
	n.pos = p
}

func hasEaten( p position , fs nodeList ) bool{

	has_haten := false
	// nodes for iteration
	var n_iter *node = fs.firstNode
	var n_previous *node = nil

	for ;n_iter !=nil;n_iter = n_iter.nextNode {
		// check if next position is one of the foods position
		if( n_iter.pos == p){
			has_haten = true
			break
		}
		n_previous = n_iter
	}

	//has eaten is true, than delete the node from the list
	if(has_haten){
		// remove food Node
		n_previous.nextNode = n_iter.nextNode
	}
	// i believe the garbage collector will collecter previous iter.

	return has_haten
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

func headIsReturning( ,) bool{
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
	// increase score for one more round
	// not sure what else
}

func (g gameState) roundIteration() bool{
	// receive position
	pos := receivePosition(g.snakeHead,g.snakeList.firstNode.nextNode.pos); // new position
	// see if has eaten, grow snake if eaten
	if ( hasEaten( pos, g.foodList ) ){
		increaseSnakeSize()
	}
	// move snake
	moveSnake(g.snakeList.firstNode,pos)
	g.snakeHead = pos

	r := false // returned bool

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
