package printer

import (
	"fmt"
)

// Hello World, é uma função de Teste para testar blibioteca.
func HelloWorld() {
	fmt.Println("Hello World")
}

func TestAscii() {
	ascBlock := string(219)
	fmt.Println(ascBlock)
}
