package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaracion de constantes se usa "const" cuando nunca cambia su valor
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion variables Enteresas
	base := 12          //Agrega valor y se inicializa sin declarar
	var altura int = 14 // Se declara, se da el tipo y se da valor
	var area int        //Se declara, se da el tipo pero no el valor

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del Cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma:", result)

	//resta
	result = y - x
	fmt.Println("Resta:", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//multiplicacion
	result = y / x
	fmt.Println("Division:", result)

	//Modulo o residio de divicion
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremental:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Tarea
	//Area Rectangulo
	baseRect := 20
	alturaRect := 50
	resultRect := baseRect * alturaRect
	fmt.Println("Area del Rectangulo:", resultRect)

	//Area Trapecio
	base1 := 10
	base2 := 20
	alturaTrap := 5
	resultTrap := (base1 + base2) * alturaTrap / 2
	fmt.Println("Area del Trapecio:", resultTrap)

	//Area de Un Circulo
	piReal := math.Pi
	radio := 10
	resultCirc := piReal * math.Pow(float64(radio), 2)
	fmt.Println("Area del Circulo:", resultCirc)

}
