package main

import "fmt"

func main() {
	//Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	// len Indica cuantos elementos hay en el array
	// cap Indica cuanta capacidad tiene el array
	fmt.Println(array, len(array), cap(array))

	//slice
	//Aqui no se indica la cantidad que va tener si no que se ingresa de una vez lo que tiene o debe tener
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el Slice
	// El primer elemento es inclusivo y el segundo no es inclusivo
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  // Imprime hasta el indice 3
	fmt.Println(slice[2:4]) // Imprime del indice 2 a 4
	fmt.Println(slice[4:])  // que imprima desde el indice 4

	//Append
	slice = append(slice, 8) //Agrega el 8 en el slice al final
	fmt.Println(slice)

	//Append nueva list (Agregar nueva lista)
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// En arrays no se pueden agragar elementos
	// En slice si se puede agregar elementos
}
