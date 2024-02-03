package main

import "fmt"

func main() {
	val,s := fn(1, 2, 3, 4, 5)
	fmt.Println(val,s)

}

func fn(values ...int) (int,string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total,"hey"
}