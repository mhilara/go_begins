package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage print a message only
func PrinrMessage(text string) {
	fmt.Println(text)
}

func prinrMessage1(text string) {
	fmt.Println(text)
}
