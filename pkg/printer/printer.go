package printer

import (
	"fmt"
)

// Hello World, é uma função de Teste para testar blibioteca.
func HelloWorld() {
	fmt.Println("Hello World")
}

//Testing function used to learn more, this was changed while developing
func TestAscii() {

	const size int = 8
	charArray := [size]int{'█','▒','╔','╚','╗','╝','═','║'} 

	for d := 0; d < size; d++ {
		fmt.Printf("%v - ",charArray[d])
		fmt.Println(string(charArray[d]))
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
}
