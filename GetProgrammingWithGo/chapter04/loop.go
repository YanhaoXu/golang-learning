package main

import "fmt"

func main() {
	count := 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
}
