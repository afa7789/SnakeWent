// // package main

// // import(
// // 	"os"
// //     "bufio"
// //     "os/exec"
// // 	"fmt"
// // )

// // func main() {
// //     // disable input buffering
// //     exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
// //     // do not display entered characters on the screen
// //     exec.Command("stty", "-F", "/dev/tty", "-echo").Run()


// //     // var b []rune = make([]rune, 1)
// //     // for {
// //     //     os.Stdin.Read(b)
// //     //     fmt.Println("I got the byte", b, "("+string(b)+") - end")
// //     // }

// //     for {
// //         // only read single characters, the rest will be ignored!!
// //         consoleReader := bufio.NewReaderSize(os.Stdin, 1)
// //         input, _ := consoleReader.ReadByte()

// //         ascii := input

// //         // ESC = 27 and Ctrl-C = 3
// //         if ascii == 27 || ascii == 3 {
// //                 fmt.Println("Exiting...")
// //                 os.Exit(0)
// //         }

// //         fmt.Println("ASCII : ", ascii)
// //     }
// // }


// package main

// import (
//     "fmt"
//     term "github.com/nsf/termbox-go"
// )

// func reset() {
// //     term.Sync() // cosmestic purpose
// }

// func main() {
//     err := term.Init()
//     if err != nil {
//         panic(err)
//     }

//     defer term.Close()

//     fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")

// keyPressListenerLoop:
//     for {
//         switch ev := term.PollEvent(); ev.Type {
//         case term.EventKey:
//             switch ev.Key {
//             case term.KeyEsc:
//                 break keyPressListenerLoop
//             case term.KeyF1:
//                 reset()
//                 fmt.Println("F1 pressed")
//             case term.KeyF2:
//                 reset()
//                 fmt.Println("F2 pressed")
//             case term.KeyF3:
//                 reset()
//                 fmt.Println("F3 pressed")
//             case term.KeyF4:
//                 reset()
//                 fmt.Println("F4 pressed")
//             case term.KeyF5:
//                 reset()
//                 fmt.Println("F5 pressed")
//             case term.KeyF6:
//                 reset()
//                 fmt.Println("F6 pressed")
//             case term.KeyF7:
//                 reset()
//                 fmt.Println("F7 pressed")
//             case term.KeyF8:
//                 reset()
//                 fmt.Println("F8 pressed")
//             case term.KeyF9:
//                 reset()
//                 fmt.Println("F9 pressed")
//             case term.KeyF10:
//                 reset()
//                 fmt.Println("F10 pressed")
//             case term.KeyF11:
//                 reset()
//                 fmt.Println("F11 pressed")
//             case term.KeyF12:
//                 reset()
//                 fmt.Println("F12 pressed")
//             case term.KeyInsert:
//                 reset()
//                 fmt.Println("Insert pressed")
//             case term.KeyDelete:
//                 reset()
//                 fmt.Println("Delete pressed")
//             case term.KeyHome:
//                 reset()
//                 fmt.Println("Home pressed")
//             case term.KeyEnd:
//                 reset()
//                 fmt.Println("End pressed")
//             case term.KeyPgup:
//                 reset()
//                 fmt.Println("Page Up pressed")
//             case term.KeyPgdn:
//                 reset()
//                 fmt.Println("Page Down pressed")
//             case term.KeyArrowUp:
//                 reset()
//                 fmt.Println("Arrow Up pressed")
//             case term.KeyArrowDown:
//                 reset()
//                 fmt.Println("Arrow Down pressed")
//             case term.KeyArrowLeft:
//                 reset()
//                 fmt.Println("Arrow Left pressed")
//             case term.KeyArrowRight:
//                 reset()
//                 fmt.Println("Arrow Right pressed")
//             case term.KeySpace:
//                 reset()
//                 fmt.Println("Space pressed")
//             case term.KeyBackspace:
//                 reset()
//                 fmt.Println("Backspace pressed")
//             case term.KeyEnter:
//                 reset()
//                 fmt.Println("Enter pressed")
//             case term.KeyTab:
//                 reset()
//                 fmt.Println("Tab pressed")

//             default:
//                 // we only want to read a single character or one key pressed event
//                 reset()
//                 fmt.Println("ASCII : ", ev.Ch)

//             }
//         case term.EventError:
//             panic(ev.Err)
//         }
//     }
// }

package main

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

func main(){
    initTerm()
    for {
        i := receiveInput(1)
        fmt.Println("aqui: ",i)
        if i==5 {
            break
        }
    }
    term.Close()
}