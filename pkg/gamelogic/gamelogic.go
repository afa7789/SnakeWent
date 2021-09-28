package gamelogic

import (
	"github.com/afa7789/SnakeWent/pkg/printer"
	"strconv"
)

type position struct {
	X, Y int
}

type node struct {
	name     string // being used for testing purposes
	pos position    // position X Y
	nextNode *node  // next node
}

type gameState struct {
	board       *[][]int
	height      int 		// board height
	width       int 		// board width
	score       int 		// game score ( number of fruits eaten + something )
	round       int 		// actual round number ( numer of iterations )
	snakeLength int 		// length of the snake ( baselength + number of eaten stuff)
	snakeHead   position	// position of the snakeHead
	snakeStart  *node 		// nodeList of snake
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

func boardTesting(){
	var board *[][]int

	board = &[][]int{{9556,9552,9552,9552,9559},{9553,183,183,183,9553},{9553,183,183,183,9553},{9553,183,183,183,9553},{9562,9552,9552,9552,9565}}

	printer.PrintSignedIntTwoDimensionsArray( board ) 
}

// First Logic
func FirstLogic() {
	printer.HelloWorld()


}
