package main

import "fmt"

func main() {
	dwarfs1 := make([]string, 0, 10)
	dwarfs1 = append(dwarfs1, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs1)

	dwarfs2 := make([]string, 10)
	dwarfs2 = append(dwarfs2, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs2)
}
