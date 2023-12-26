package main

import "fmt"

func main() {
	//Decalaraciones
	helloMessage := "Hello"
	worldMessage := "World"

	//Println //Imprime y agrega un enter o salto de linea
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos) //Esepicida el tipo de dato
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos) //con v puede ser cualquiera

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos //para saber el tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
