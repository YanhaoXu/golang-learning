package main

import "fmt"

func main() {
	var administrator *string

	scolese := "Christoper J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)
}
