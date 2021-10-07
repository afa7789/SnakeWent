package game

// modificar para usar o printer.

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

const ( // this are used to drawn the board
	arrowUp    = 1
	arrowDown  = 2
	arrowLeft  = 3
	arrowRight = 4
	esc        = 5
	rKey       = 6
)

// clear the terminal
func reset() {
	term.Sync() // cosmestic purpose
}

// initialize the usage of TermBox
func initTerm() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
}

// Uses Termbox Library to receive the keyboard inputs arrowkeys to designate the new direction
func receiveInput(not_this_dir int) int {

	fmt.Println("Enter arrow keys, r to restart, or press esc button to quit")
keyPressListenerLoop:
	for {
		var key_return int = 0
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				key_return = esc
				break keyPressListenerLoop
			case term.KeyArrowUp:
				key_return = arrowUp
				break
			case term.KeyArrowDown:

				key_return = arrowDown
				break
			case term.KeyArrowLeft:
				key_return = arrowLeft
				break
			case term.KeyArrowRight:
				key_return = arrowRight
				break
			default:
				if ev.Ch == 114 {
					key_return = rKey
					break
				} else {
					fmt.Println("You pressed an invalid key")
					fmt.Println("Enter arrow keys, r to restart, or press esc button to quit")
				}
			}
			break
		case term.EventError:
			panic(ev.Err)
		}

		if not_this_dir == key_return {
			fmt.Println("You can't go this direction")
			fmt.Println("Enter other arrow keys")
		}

		if not_this_dir != key_return && key_return != 0 {
			return key_return
		}
	}
	return 5
}
