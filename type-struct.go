package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	mycar := car{brand: "Ford", year: 2022}
	fmt.Println(mycar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Aston Martin"
	fmt.Println(otherCar)
}
