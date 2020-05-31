package main

import (
	"fmt"
	"os"
)

func hoge() {
	defer fmt.Println("world hoge")
	fmt.Println("hello hoge")
}

func main() {
	hoge()
	defer fmt.Println("WORLD")
	fmt.Println("HELLO")

	// stacking defers
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	file, _:= os.Open("./map.go")
	defer file.Close()
	data := make([]byte,100)
	file.Read(data)
	fmt.Println(string(data))
	fmt.Println()

}


