// Modificadores de acceso funciones y structs
package main

import (
	"fmt"

	pk "/go_begins/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Aston Martin"
	fmt.Println(myCar)

	var myOtherCar pk.carPrivate
	fmt.Println(myOtherCar)

	pk.PrinrMessage("Hola mundo")
	pk.prinrMessage("Hello private")
}
