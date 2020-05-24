package main

import "fmt"
/*
	関数のsyntax確認
 */
func main() {
	result := add(10,20)
	fmt.Println(result)

	result1, result2 := addReturnTwoVariable(10,20)
	fmt.Println(result1, result2)

	result3 := calc(100,20)
	fmt.Println(result3)

	// 関数内に関数を定義できる
	f := func(str string) {
		fmt.Println("inner func", str)
	}
	f("test")
	// 上記と同じ結果
	func(str string) {
		fmt.Println("inner func", str)
	}("test")

}

func add(x, y int) int {
	return x + y
}

func addReturnTwoVariable(x, y int) (int, int) {
	return x + y, x -y
}

func calc(price, item int) (result int) {
	result = price * item
	// ネイキッドリターン
	return
}