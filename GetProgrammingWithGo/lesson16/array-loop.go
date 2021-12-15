package main

import "fmt"

func main() {
	fmt.Println("####### array-loop ########")
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
	fmt.Println("####### array-range ########")
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

}
