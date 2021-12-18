package main

import "fmt"

func main() {
	temerature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temerature["Earth"]
	fmt.Printf("On average the Earth is %v°C.\n", temp)
	temerature["Earth"] = 16
	temerature["Venus"] = 464
	fmt.Println(temerature)

	moon := temerature["Moon"]
	fmt.Println(moon)

	if moon, ok := temerature["Moon"]; ok {
		fmt.Printf("On average the moon is %v°C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

}
