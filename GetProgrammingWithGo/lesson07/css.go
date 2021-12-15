package main

import "fmt"

func main() {
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("%x %x %x", red, green, blue)
	fmt.Println()
	fmt.Printf("%X %X %X", red, green, blue)
	fmt.Println()
	fmt.Printf("color: #%02X%02X%02X", red, green, blue)
}
