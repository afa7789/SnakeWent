package game

import (
	"github.com/afa7789/SnakeWent/pkg/printer"
	"strconv"
)

const ( // this are used to drawn the board
	rectangle = 9608		//	9608 - █ , place occupied
	meshRectangle = 9618	//	9618 - ▒ , the conflict (die drawing)
	dot = 183 				//	0183 - · , empty space
	diamondEmpty = 9671		// 	9671 - ◇ , stuff to eat
	// borders
	cornerUpRight = 9559	//	9559 - ╗ 
	cornerDownRight = 9565	//	9565 - ╝
	cornerUpLeft = 9556 	//	9556 - ╔
	cornerDownLeft = 9562 	//	9562 - ╚
	cornerHorizontal = 9552 //	9552 - ═
	cornerVertical = 9553  	//	9553 - ║
)

type position struct {
	X, Y int
}

type node struct {
	name     string // being used for testing purposes
	pos position    // position X Y
	nextNode *node  // next node
}

type nodeList struct{
	firstNode *node // first node for quicker access
	lastNode  *node  // last node for quicker access
}

type gameState struct {
	board       *[][]int		// board 2d array
	height      int 			// board height
	width       int 			// board width
	score       int 			// game score ( number of fruits eaten + something )
	round       int 			// actual round number ( numer of iterations )
	snakeLength int 			// length of the snake ( baselength + number of eaten stuff)
	snakeHead   position		// position of the snakeHead
	snakeStart  nodeList 		// nodeList of snake
	food nodeList				// nodeList of food
}

// GAME STATE
// print game state boards
func (g *gameState) Print(){
	printer.PrintSignedIntTwoDimensionsArray( g.board )
}

///NODE FUNCTIONS
// print node, to see how it is
func (n *node) Print(){
	stringForPrint := "Node print"
	stringForPrint += "Name: " + n.name + "\n"
	stringForPrint += "X: " + strconv.Itoa(n.pos.X) + "\n"
	stringForPrint += "Y: " + strconv.Itoa(n.pos.Y) + "\n"

	printer.PrintString(stringForPrint)
}

// create Node, return a new pointer to node
func createNode(s string, x,y int, n *node) *node{
	new_n := &node{
		name: s,
		pos: position{
			X: x,
			Y: y,
		},
		nextNode: n,
	}
	return new_n
}

// create nodeList, return new nodeList
func createNodeList(n1 *node,n2 *node) nodeList{
	return nodeList{
		firstNode:n1,
		lastNode:n2,
	}
}

// delete a node from a food node list.
func deleteNode(nl nodeList,n *node){
	// receives node, 
	var n_iter = nl.firstNode

	var n_previous *node = nil

	for ;n_iter !=nil;n_iter = n_iter.nextNode {
		// comparar node
		if( n_iter == n){
			break
		}
		n_previous = n_iter
	}

	n_previous.nextNode = n_iter.nextNode
}

// change nextNode, change the next node of a node, and returns the previous node it had pointed.
func (n *node) changeNextNode(w *node) *node{
	var previousNode *node = n.nextNode;
	n.nextNode = w;
	return previousNode;
}