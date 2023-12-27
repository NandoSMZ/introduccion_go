package main

import "fmt"

func main() {
	//Defer hace que se ejecute como la ultima linea
	// Es decir que primero imprime "Mundo" y luego "Hola"
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue y Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue, sirve cuando hay errores
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break
		if i == 6 {
			fmt.Println("Es 6 BREAK")
			break
		}
	}
}
