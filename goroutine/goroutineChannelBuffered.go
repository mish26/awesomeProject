package main

import "fmt"

func main() {
	chanInt := make(chan int, 2)
	chanInt <- 100
	fmt.Println(len(chanInt))
	chanInt <- 100
	fmt.Println(len(chanInt))

	v := <- chanInt
	fmt.Println(v)

	chanInt <- 200
	fmt.Println(len(chanInt))

	// closeしないと、forがエラーになる
	// 2つしかないチャネルに対し3つめを取ろうとしてしまう
	close(chanInt)

	for c := range chanInt {
		fmt.Println(c)
	}
}



