package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "For", year: 2020}
	fmt.Println(myCar)

	//Otra manera de instanciar
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2022
	fmt.Println(otherCar)
}
