
import (
    "fmt"
    term "github.com/nsf/termbox-go"
)

const ( // this are used to drawn the board
	ARROW_UP = 1
	ARROW_DOWN = 2
	ARROW_LEFT = 3
	ARROW_RIGHT = 4
	ESC = 5
	R_KEY = 6 
)

func reset() {
	term.Sync() // cosmestic purpose
}

func initTerm(){
	err := term.Init()
	if err != nil {
		panic(err)
	}
}

func receiveInput(not_this_dir int) int {

	fmt.Println("Enter arrow keys, r to restart, or press ESC button to quit")
    keyPressListenerLoop: for {
		var key_return int = 0
		switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
					case term.KeyEsc:
    					key_return = ESC
                        break keyPressListenerLoop
					case term.KeyArrowUp:
						key_return = ARROW_UP
						break			
					case term.KeyArrowDown:
                        
						key_return = ARROW_DOWN
						break					
					case term.KeyArrowLeft:
						key_return = ARROW_LEFT
						break				
					case term.KeyArrowRight:
						key_return = ARROW_RIGHT
						break
					default:
						if ev.Ch == 114{
                            key_return = R_KEY
							break
						}else{
                            fmt.Println("You pressed an invalid key")
							fmt.Println("Enter arrow keys, r to restart, or press ESC button to quit")
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

		if not_this_dir != key_return && key_return != 0{
			return key_return
		}else{

		}
	}
    return 5
}
