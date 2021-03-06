package game

import (
	"strconv"

	"github.com/afa7789/SnakeWent/pkg/printer"
)

const ( // this are used to drawn the board
	rectangle     = 9608 //	9608 - █ , place occupied
	meshRectangle = 9618 //	9618 - ▒ , the conflict (die drawing)
	dot           = 183  //	0183 - · , empty space
	diamondEmpty  = 9671 // 	9671 - ◇ , stuff to eat
	// borders
	cornerUpRight    = 9559 //	9559 - ╗
	cornerDownRight  = 9565 //	9565 - ╝
	cornerUpLeft     = 9556 //	9556 - ╔
	cornerDownLeft   = 9562 //	9562 - ╚
	cornerHorizontal = 9552 //	9552 - ═
	cornerVertical   = 9553 //	9553 - ║
)

// x,y pos struct
type position struct {
	X, Y int
}

// node struct, with name, pos and next node
type node struct {
	name     string   // being used for testing purposes
	pos      position // position X Y
	nextNode *node    // next node
}

// last node and first node
type nodeList struct {
	firstNode *node // first node for quicker access
	lastNode  *node // last node for quicker access
}

//gamestate variables and others
type gameState struct {
	board       *[][]int // board 2d array
	height      int      // board height
	width       int      // board width
	score       *int     // game score ( number of fruits eaten + something )
	round       *int     // actual round number ( numer of iterations )
	snakeLength *int     // length of the snake ( baselength + number of eaten stuff)
	snakeHead   position // position of the snakeHead
	snakeList   nodeList // nodeList of snake
	foodList    nodeList // nodeList of food
}

// GAME STATE
// print game state boards
func (g *gameState) Print(b bool) {
	stringForPrint := "Game State\n"
	// stringForPrint += "Height: " + strconv.Itoa(g.height) + " - Width: " + strconv.Itoa(g.width) + "\n"
	stringForPrint += "\tScore: " + strconv.Itoa(*g.score) + " - Round: " + strconv.Itoa(*g.round) + "\n"
	printer.PrintString(stringForPrint)
	if b {
		printer.PrintSignedIntTwoDimensionsArray(g.board)
	}
}

///NODE FUNCTIONS
// print node, to see how it is
func (n *node) Print() {
	stringForPrint := "Node print - "
	stringForPrint += n.name + "\n"
	stringForPrint += "\t X: " + strconv.Itoa(n.pos.X) + "\n"
	stringForPrint += "\t Y: " + strconv.Itoa(n.pos.Y) + "\n\n"

	printer.PrintString(stringForPrint)
}

func (p position) Print() {
	printer.PrintString("[" + intToString(p.X) + "," + intToString(p.Y) + "]")
}

// nodeListTesting is created to make sure list
func (nl nodeList) Print() {
	n_iter := nl.firstNode
	for ; n_iter != nil; n_iter = n_iter.nextNode {
		n_iter.Print()
	}
}

// create Node, return a new pointer to node
func createNode(s string, x, y int, n *node) *node {
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
func createNodeList(n1 *node, n2 *node) nodeList {
	return nodeList{
		firstNode: n1,
		lastNode:  n2,
	}
}

// this returns an zeroed gameState
func createGameStateZero() gameState {

	nl1 := createNodeList(nil, nil)
	nl2 := createNodeList(nil, nil)

	return gameState{
		board:       nil,
		height:      0,
		width:       0,
		score:       nil,
		round:       nil,
		snakeLength: nil,
		snakeHead: position{
			X: 0,
			Y: 0,
		},
		snakeList: nl1,
		foodList:  nl2,
	}
}

func createGameState(h, w int) gameState {
	initTerm()
	g := createGameStateZero()
	g.height = h
	g.width = w
	// creating board
	board := make([][]int, h+2)
	for i := range board {
		board[i] = make([]int, w+2)
	}
	round := 0
	score := 0
	length := 3
	g.round = &round
	g.score = &score
	g.snakeLength = &length
	var i int = (w + 1) / 2
	var j int = (h + 1) / 2

	var n3 = createNode("Snake3", i, j, nil)
	var n2 = createNode("Snake2", i+1, j, n3)
	var n1 = createNode("Snake1", i+1, j+1, n2)
	var nl = createNodeList(n1, n3)
	g.snakeList = nl
	g.snakeHead = g.snakeList.firstNode.pos
	g.board = &board
	fillBorders(g.board, g.width, g.height)
	fillAll(g)
	return g
}

func intToString(i int) string {
	return strconv.Itoa(i)
}
