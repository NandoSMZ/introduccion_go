package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es Palindomo")
	} else {
		fmt.Println("No es Palindromo por que la palabra es:", textReverse)
	}
}

func main() {
	slice := []string{"Hola", "Mundo", "Que Hace"}

	//Para traer indice y valor
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// Solo necesito el valor
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// Solo indice
	for i := range slice {
		fmt.Println(i)
	}

	var palabra string
	fmt.Println("Ingrese la palabra")
	fmt.Scan(&palabra)
	palabraToMinus := strings.ToLower(palabra)
	fmt.Println("Llamando la funcion para la palabra:", palabra)
	isPalindromo(palabraToMinus)
}
