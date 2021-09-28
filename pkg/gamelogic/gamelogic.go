package gamelogic

import (
	"github.com/afa7789/SnakeWent/pkg/printer"
)

type position struct {
	X, Y int
}

type node struct {
	name     string
	pos position
	nextNode *node
}

type gameState struct {
	board       *[][]int
	height      int //board height
	width       int //board width
	score       int //game score ( number of fruits eaten + something )
	round       int //actual round number ( numer of iterations )
	snakeLength int // length of the snake ( baselength + number of eaten stuff)
	snakeHead   position
	snakeStart  *node // position of the snakeHead
}

func (n *node) Print(){
	stringForPrint := "Node print"
	stringForPrint += "Name: " + n.name + "\n"
	stringForPrint += "X: " + string(n.pos.X) + "\n"
	stringForPrint += "Y: " + string(n.pos.Y) + "\n"

	printer.PrintString(stringForPrint)
}


func (g *gameState) Print(){

	printer.PrintSignedIntTwoDimensionsArray( g.board )

}

//
/*	
	9608 - █
	9618 - ▒
	9556 - ╔
	9562 - ╚
	9559 - ╗
	9565 - ╝
	9552 - ═
	9553 - ║
*/

// First Logic
func FirstLogic() {
	printer.HelloWorld()

	var board *[][]int

	board = &[][]int{{9556,9552,9559},{9553,183,9553},{9562,9552,9565}}

	var n = node{
		name: "Arthur",
		pos: position{
			X: 0,
			Y: 0,
		},
		nextNode: nil,
	} 

	n.Print();

	printer.PrintSignedIntTwoDimensionsArray( board ) 
}
