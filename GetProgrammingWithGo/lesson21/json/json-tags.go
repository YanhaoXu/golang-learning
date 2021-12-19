package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}
