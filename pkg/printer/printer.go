package printer

import (
	"fmt"
)

// Hello World, é uma função de Teste para testar blibioteca.
func HelloWorld() {
	fmt.Println("This is Snake Went a dubious quality game\n")
}

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
//Testing function used to learn more, this was changed while developing
func TestAscii() {

	const size int = 8
	charArray := [size]int{'█', '▒', '╔', '╚', '╗', '╝', '═', '║'}

	for d := 0; d < size; d++ {
		fmt.Printf("%v - ", charArray[d])
		fmt.Println(string(charArray[d]))
	}
}

func PrintString( str string ){
	fmt.Printf(str);
}

func PrintSignedIntTwoDimensionsArray( sample * [][]int  ){
    for _, row := range *sample {
        for _, val := range row {
            fmt.Print(val)
        }
		fmt.Printf("\n")
    }
}
