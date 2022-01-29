package main

import "fmt"

func main() {
	// declaracion de variales
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Milton"
	novia := 1
	fmt.Printf("%s tiene %d novia bien bonita\n", nombre, novia)
	fmt.Printf("%v tiene %v novia bien bonita\n", nombre, novia)

	// Sprintf
	message := fmt.Sprintf("%s tiene %d novia bien bonita", nombre, novia)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("novia: %T\n", novia)

	// Hola mundo a mi estilo
	fmt.Println("Hola mundo")

}
