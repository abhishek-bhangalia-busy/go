package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	var y *float64 = &x

	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(y).Elem())

	fmt.Println()

	fmt.Println(reflect.ValueOf(y))
	fmt.Println(reflect.ValueOf(y).Interface())

	fmt.Println()

	fmt.Println(reflect.ValueOf(y).Elem())
	fmt.Println(reflect.ValueOf(y).Elem().Interface())
	fmt.Println(reflect.ValueOf(*y))

	fmt.Println()

	// You can also convert it to the underlying type explicitly
	xValue := reflect.ValueOf(y).Elem().Interface().(float64)
	fmt.Println("Value of x:", xValue)

	fmt.Println()

	reflect.ValueOf(y).Elem().SetFloat(7.8)
	fmt.Println("Original value of x changed, x=", x)
}
