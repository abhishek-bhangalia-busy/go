package main

import (
	"fmt"
)

// func main() {

// 	var i int
// 	var j string

// 	fmt.Scan(&i, &j)
// 	fmt.Println(i, j)

// }

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n');
	fmt.Printf("%T",input)
	fmt.Println("\nrating is: ",input)

// newrating, err := strconv.ParseFloat(input,64)

	newrating, err := strconv.ParseFloat(strings.TrimSpace(input),64)
	if(err!=nil){
		fmt.Println(err)
	}else{
		fmt.Println(newrating+1)
	}

}
