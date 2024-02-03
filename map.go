package main

import (
	"fmt"
)

func main() {

	mp := make(map[string]string)

	mp["a"] = "apple"
	mp["b"] = "ball"
	mp["c"] = "cat"

	fmt.Println(mp)

	delete(mp, "b")
	if _, ok := mp["d"]; ok {
		fmt.Println(mp,"sdfsd")
	}
	
	fmt.Println(mp["d"])

	for key, val := range mp {
		fmt.Println("for key", key, "the value is", val)
	}

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Accessing values
	fmt.Println("Value for key 'one':", myMap["one"])

}
