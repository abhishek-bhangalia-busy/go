package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func main() {

	score := []int{}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Score is 1")
		score = append(score, 1)
		defer wg.Done()
	}(&wg, &mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Score is 2")
		score = append(score, 2)
		defer wg.Done()
	}(&wg, &mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Score is 3")
		score = append(score, 3)
		defer wg.Done()
	}(&wg, &mut)

	wg.Wait()

	fmt.Println(score)
}
