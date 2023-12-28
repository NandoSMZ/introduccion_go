package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Personalizar el output de STRUCT personalizado
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "MSI", disk: 200}
	fmt.Println(myPC)
}
