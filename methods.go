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

	User1.changeName()
	fmt.Println(User1.Name)
}

func (usr1 User) changeName (){

	usr1.Name="Bhardwaj"
	fmt.Println(usr1.Name)
}
