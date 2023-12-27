package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	//With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es Verdad")
	}

	//With or
	if valor1 == 1 || valor2 == 2 {
		fmt.Println("Es Verdad")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	//Si no existe error va ser nil
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)

	//Reto de saber si uyn numero es par y de verificar password
	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}
}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}
