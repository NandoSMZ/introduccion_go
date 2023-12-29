package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2) // canal con capacidad maxima de 2
	c <- "mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c)) // len para saber la cantidad de datos y cap para la capacidad del canal

	//Range y Close

	close(c) // Va cerrar el canal, es decir que no va recibir mas datos el cnal

	for message := range c {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string) //Canal dinamico
	email2 := make(chan string)
	go message("Mensaje 1 Email1", email1)
	go message("Mensaje 2 Email2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email Recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email Recibido de email2", m2)
		}
	}
}
