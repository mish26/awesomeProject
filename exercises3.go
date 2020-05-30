package main

import (
	"fmt"
)

type Vertex4 struct{
	X, Y int
}

func (v Vertex4) Plus() int {
	return v.X + v.Y
}

func main(){
	v := Vertex4{3, 4}
	fmt.Println(v.Plus())
}
