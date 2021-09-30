package gamelogic

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
	snakeStart  *nodeList 		// nodeList of snake
	food *nodeList				// nodeList of food
}

func (n *node) Print(){
	stringForPrint := "Node print"
	stringForPrint += "Name: " + n.name + "\n"
	stringForPrint += "X: " + strconv.Itoa(n.pos.X) + "\n"
	stringForPrint += "Y: " + strconv.Itoa(n.pos.Y) + "\n"

	printer.PrintString(stringForPrint)
}

func (g *gameState) Print(){

	printer.PrintSignedIntTwoDimensionsArray( g.board )

}

func createNode(s string, x,y int, n &node){
	return *node{
		name: s,
		pos: position{
			X: x,
			Y: y,
		},
		nextNode: n,
	}
}

func createNodeList(n1 *node,n2 *node){
	return nodeList{
		firstNode:n1,
		lastNode:n2,
	}
}

func (n *node) changeNextNode(w *node){
	var previousNode *node = n.nextNode;
	n.nextNode = w;
	return previousNode;
}

func nodeTesting(){
	var n = node{
		name: "Arthur",
		pos: position{
			X: 0,
			Y: 0,
		},
		nextNode: nil,
	} 

	n.Print();
}

// nodeListTesting is created to make sure list
func nodeListTesting(){
	var n1 = createNode("Snake1",0,0,nil)
	var n2 = createNode("Snake2",0,0,nil)
	n1.changeNextNode(&n2)
	var nl = createNodeList(&n1,&n2)
}

// fillBorders changes int from board state to empty board
func fillBorders(){
   // fill borders , 
}

// fillEmptyDots, changes int from board state to empty board
func fillEmptyDots(){
	// fill everything that's not border
}

// delete a node from a food node list.
func deleteFood(){
	// receives node, 
	// not sure ask question
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

func boardTesting(){
	var board *[][]int

	board = &[][]int{{9556,9552,9552,9552,9559},{9553,183,183,183,9553},{9553,183,183,183,9553},{9553,183,183,183,9553},{9562,9552,9552,9552,9565}}

	printer.PrintSignedIntTwoDimensionsArray( board ) 
}

// First Logic
func FirstLogic() {
	printer.HelloWorld()


}
