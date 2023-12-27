package main

import "fmt"

func normalFuntion(message string) {
	fmt.Println(message)
}

func tripeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

// Recibe un INT y devuelve un INT
func returnValue(a int) int {
	return a * 2
}

// Recibe 1 valor A, pero retorna dos valores c y d
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuntion("Hola Mundo")
	tripeArgument(1, 2, "Hola Go")

	value := returnValue(2)
	fmt.Println(value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Value1 y Value2:", value1, value2)

	// SI necesito un solo valor, entonces...
	value1Only, _ := doubleReturn(4)
	fmt.Println("value1Only:", value1Only)
}
