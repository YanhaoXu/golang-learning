package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]

	allPlanets := planets[:]

	fmt.Println(terrestrial)
	fmt.Println(gasGiants)
	fmt.Println(iceGiants)
	fmt.Println(allPlanets)
}
