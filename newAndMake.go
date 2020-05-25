package main

import "fmt"

func main() {

	// 変数に何も入れず、メモリにポインターの領域(アドレス)を確保したい
	var pointer *int = new(int)
	fmt.Println(pointer, *pointer)
	*pointer++
	fmt.Println(*pointer)

	var pointer2 *int
	fmt.Println(pointer2)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// メモリ領域を確保していない状態で、状態を変更することはできない。panicになる。
	// *pointer2++
	fmt.Println(pointer2)


}

