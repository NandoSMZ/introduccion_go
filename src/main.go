package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer map

	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar Valor
	// con el ok, vamos a saber si la variable existe y sera booleano
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
