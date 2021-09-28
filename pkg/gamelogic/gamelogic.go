package gamelogic

import (
	"../printer"
)

type position struct{
	X,Y int
}

type gameState struct{
	board *[][]int
	score int
	round int
	snakeLength int
	height int
	width int
	snakeHead position
}

type snakeNode struct{
	snakePos position
	nextNode *snakeNode
}

// First Logic
func FirstLogic() {
	printer.HelloWorld()
}