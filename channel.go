package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int, 2) //2 in buffer can go

	wg.Add(2)

	go func(wg *sync.WaitGroup, ch chan int) {
		fmt.Println(<-ch)
		defer wg.Done()
	}(wg, ch)

	go func(wg *sync.WaitGroup, ch chan int) {
		ch <- 4
		ch <- 6
		defer wg.Done()
	}(wg, ch)

	wg.Wait()
}
