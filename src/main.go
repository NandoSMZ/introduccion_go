package main

import "fmt"

func say(text string, c chan string) {
	c <- text //Vamos a ingresar el text al chanel
}

func main() {
	//creacion de chanel especificando cuantas goroutines van a pasar en esste caso "1" si se deja basio sera default o dinamico
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c) //ejecutamos la funcion indicando el chanel c

	fmt.Println(<-c) // Sacamos el dato que esta en el channel
}
