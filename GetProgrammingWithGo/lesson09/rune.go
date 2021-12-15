package main

import "fmt"

func main() {
	// type byte = int8
	// type rune = int32

	var pi rune = 960
	var alpha rune = 940

	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)

	grade := 'A'
	fmt.Printf("%c\n", grade)
}
