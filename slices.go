package main

import (
	"fmt"
	"sort"
)

func main() {

	var sl =[]int{5,3,9,1}         //slice declaring without any memory value at start
	fmt.Println(sl)

	sl = append(sl, 10,22,43)
	fmt.Println(sl)

	sl=append(sl[1:] )
	fmt.Println(sl)



	sl2 := make([]int,4)

	sl2[0]=9
	sl2[1]=4
	sl2[2]=0
	sl2[3]=2

 	fmt.Println("\nThis is slice 2")
	fmt.Println(sl2)

	sl2 = append(sl2, 200,100)
	fmt.Println(sl2)

	sort.Ints(sl2)
	fmt.Println(sl2)

	fmt.Println(sort.IntsAreSorted(sl2))

	//Remove index=2 from slice

	var ind int=2
	
	sl2= append(sl2[:ind],sl2[ind+1:]...)        
	fmt.Println(sl2)

	delete(sl2,2)
	fmt.Println(sl2)
}