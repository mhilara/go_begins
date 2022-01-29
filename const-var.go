package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero  values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Retos
	// -Rectangulo, trapecio y de un circulo

	// Formula rectangulo Area=largo×ancho
	const largoRectangulo = 50
	const anchoRectangulo = 25
	areaRectangulo := largoRectangulo * anchoRectangulo
	fmt.Println("Area rectangulo:", areaRectangulo)

	// Formula trapecio area = h*(a+b/2)
	const base1 = 10
	const base2 = 5
	const alturatra = 7
	areaTrapecio := alturatra * ((base1 + base2) / 2)
	fmt.Println("Area trapecio:", areaTrapecio)

	// Formula area circulo = area = pi*r2
	const valorpi = 3.1416
	const radio = 2
	areaCirculo := valorpi * math.Pow(radio, 2)
	fmt.Println("Area Circulo:", areaCirculo)

}
