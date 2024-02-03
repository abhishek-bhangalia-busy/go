package main

import "fmt"

type Player interface {
	Style() int
	Name() string
}

type Messi struct {
	Pass int
	Sui  int
}

type CR7 struct {
	Pass    int
	Dribble int
	Sui     int
}

func (m Messi) Style() int{
	ts := m.Pass*m.Sui
	// fmt.Printf("Total style is: %v \n", ts)
	return ts
}
func (c CR7) Style() int {
	ts := c.Pass*c.Dribble*c.Sui
	// fmt.Printf("Total style is: %v \n", ts)
	return ts
}

func (m Messi) Name() string{
	return "Messi"
}
func (c CR7) Name() string{
	return "CR7"
}


func main() {

	var m Messi
	m.Pass=10
	m.Sui=9

	var c =CR7{
		Pass:10,
		Dribble:10,
		Sui:10,
	}

	var sl= []Player{m,c}

	for i:=0;i<len(sl);i++{
		// sl[i].Style()

		fmt.Printf("%v style is %v \n",sl[i].Name(),sl[i].Style())
	}

}