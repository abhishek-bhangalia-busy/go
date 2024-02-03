package main

import (
	"fmt"
	"math/rand"
)

type User struct {
	Name    string
	Age     int
	Married bool
}

func main() {

	User1 := User{"Vaibhav", 22, false}
	fmt.Printf("Details are %+v\n", User1)
	fmt.Printf("Name of the user is %v\n", User1.Name)

	if num := 11; num < 10 {
		fmt.Println("less")
	} else {
		fmt.Println("greater")
	}

	var rval int = rand.Intn(6)+1
	fmt.Println(rval)

}
