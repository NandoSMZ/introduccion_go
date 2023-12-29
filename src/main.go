package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // Acomula GO routins y las libera poco a poco

	fmt.Println("Hello") // Imprime primero
	wg.Add(1)            // Se agrega la go rotin al waitgroup

	go say("World", &wg) // Se ejecuta de forma concurrente

	wg.Wait() // Espere hasta que todas las GO Rountins se finalicen

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	// Se agrega el contador par que el WORLD alcance a salir

	time.Sleep(time.Second * 1)

}
