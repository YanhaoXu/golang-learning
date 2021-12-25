package main

import "fmt"

func main() {
	superpower := &[3]string{"flight", "invisibility", "super strength"}

	fmt.Println(superpower[0])
	fmt.Println(superpower[1:2])
}
