package gamelogic

import (
	"github.com/afa7789/SnakeWent/pkg/printer"
)

type position struct {
	X, Y int
}

type snakeNode struct {
	name     string
	snakePos position
	nextNode *snakeNode
}

type gameState struct {
	board       *[][]int
	height      int //board height
	width       int //board width
	score       int //game score ( number of fruits eaten + something )
	round       int //actual round number ( numer of iterations )
	snakeLength int // length of the snake ( baselength + number of eaten stuff)
	snakeHead   position
	snakeStart  *snakeNode // position of the snakeHead
}

func snakeNodePrint(node *snakeNode){
	
	stringForPrint := ""
	stringForPrint += "Name: " + node.name + "\n"
	stringForPrint += "X: " + node.name + "\n"
	stringForPrint += "Y: " + node.name + "\n"

	printer.PrintString(stringForPrint)
}

// First Logic
func FirstLogic() {
	printer.HelloWorld()
	snakeNodePrint()
}
