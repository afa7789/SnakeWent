package game

import "github.com/afa7789/SnakeWent/pkg/printer"

func nodeTesting() {
	var n = node{
		name: "Arthur",
		pos: position{
			X: 0,
			Y: 0,
		},
		nextNode: nil,
	}

	n.Print()
}

// nodeListTesting is created to make sure list
func nodeListTesting() {
	var n3 = createNode("Snake3", 0, 0, nil)
	var n2 = createNode("Snake2", 0, 0, n3)
	var n1 = createNode("Snake1", 0, 0, n2)

	var nl = createNodeList(n1, n3)

	n_iter := nl.firstNode

	for ; n_iter != nil; n_iter = n_iter.nextNode {
		n_iter.Print()
	}
}

// check if boardPrinting is working ok
func boardTesting() {
	var board *[][]int

	board = &[][]int{{9556, 9552, 9552, 9552, 9559}, {9553, 183, 183, 183, 9553}, {9553, 183, 183, 183, 9553}, {9553, 183, 183, 183, 9553}, {9562, 9552, 9552, 9552, 9565}}

	printer.PrintSignedIntTwoDimensionsArray(board)
}

// First Logic
func FirstLogic() {
	printer.HelloWorld()

}
