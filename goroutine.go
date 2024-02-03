package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup


func main() {
	sl := []int{1, 2, 3, 4, 5}
	

	for _, val := range sl {
		go fn(val)
		wg.Add(1)
	}


	wg.Wait()
}

func fn(v int) {

	defer wg.Done()
	fmt.Println(v)

}