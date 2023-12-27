package main

import "fmt"

func main() {
	//Se define la variable y luego se hace el parceo de esa variable o variable a evaluar
	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es Impar")
	}

	//sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No Condicion")
	}
}
