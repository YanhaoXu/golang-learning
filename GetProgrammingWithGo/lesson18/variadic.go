package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorld := make([]string, len(worlds))
	for i := range worlds {
		newWorld[i] = prefix + " " + worlds[i]
	}
	return newWorld
}

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)

	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}
