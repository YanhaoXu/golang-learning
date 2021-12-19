package main

import "fmt"

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)

}

func main() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
}
