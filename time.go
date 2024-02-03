package main

import (
	"fmt"
	"time"
)

func main() {

	curdate := time.Now()
	fmt.Println(curdate.Format( "02/01/2006 Monday"))
}