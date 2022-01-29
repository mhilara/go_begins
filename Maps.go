package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Milton"] = 30
	m["Hilara"] = 15

	fmt.Println(m)

	// Recorre map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["Milton"]
	fmt.Println(value, ok)
}

