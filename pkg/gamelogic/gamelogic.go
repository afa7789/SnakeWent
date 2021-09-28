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
	stringForPrint += "X: " + n.name + "\n"
	stringForPrint += "Y: " + n.name + "\n"

	printer.PrintString(stringForPrint)
}


func (g *gameState) Print(){

	printer.PrintSignedIntTwoDimensionsArray( g.board )

}

// First Logic
func FirstLogic() {
	printer.HelloWorld()

	var board *[][]int

	board = &[][]int{{1,22,2},{2,},{3,}}

	var n = node{
		name: "Arthur",
		pos: position{
			X:0,
			Y:0,
		},
		nextNode: nil,
	} 

	n.Print();

	printer.PrintSignedIntTwoDimensionsArray( board ) 
}
