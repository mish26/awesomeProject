package main

import (
	"fmt"
)

type Vertex5 struct{
	X, Y int
}

func (v Vertex5) String() string {
//	return fmt.Sprintf("X is %v!", v.X ) + fmt.Sprintf(" Y is %v!", v.Y)
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}

func main(){
	v := Vertex5{3, 4}
	fmt.Println(v)
}
