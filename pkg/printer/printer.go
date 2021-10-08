package printer

import (
	"fmt"
)

// Hello World, é uma função de Teste para testar blibioteca.
func HelloWorld() {
	fmt.Println("This is Snake Went a dubious quality game\n")
}

func checkCharacter(i int) {
	fmt.Printf("%v - %s \n", i, string(i))
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
	9671 - ◇
	183 - ·
	9708 - ◬
*/
//Testing function used to learn more, this was changed while developing
func TestAscii() {

	const size int = 8
	charArray := [size]int{'█', '▒', '╔', '╚', '╗', '╝', '═', '║'}

	for d := 0; d < size; d++ {
		checkCharacter(charArray[d])
		// fmt.Printf("%v - ", charArray[d])
		// fmt.Println(string(charArray[d]))
	}

	// fmt.Printf("%v",int('◬'));
	// checkCharacter('◬')
}

func PrintString(str string) {
	fmt.Printf("\t" + str)
}

func PrintSignedIntTwoDimensionsArray(sample *[][]int) {
	for _, row := range *sample {
		fmt.Print("\t")
		for _, val := range row {
			fmt.Print(string(val), "")
		}
		fmt.Printf("\n")
	}
}
