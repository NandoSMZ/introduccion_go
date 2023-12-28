package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

// COn el puntero ingreo directamente al valor en memoria
func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a //obtiene direccion en memoria

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // Obtiene el valor que apunta esa dir de memoria

	myPC := pc{ram: 16, disk: 200, brand: "MSI"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC)
}
