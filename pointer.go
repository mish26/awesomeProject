package main

import "fmt"

func sampleFunc(x *int)  {
	fmt.Println("sampleFunc:", x, *x)
	// xのメモリアドレスに、1を代入する
	*x = 1
	fmt.Println("sampleFunc:", x, *x)
}
func main() {
	var num int = 1000
	fmt.Println("num: ", num)
	// 1000を入れたメモリのアドレスを出力
	fmt.Println("&num: ",&num)

	// integerのポイント型
	// &はメモリのアドレスを指している。int型だとエラーになる。
	var point *int = &num
	fmt.Println("point:", point,",point*:", *point)

	sampleFunc(&num)

	// sampleFunc内で、指定のメモリアドレスに値を代入すると上書きされる
	// sampleFuncにメモリのアドレスを代入しているので、
	// phpの参照渡しと一緒
	fmt.Println("sampleFunc後:", num)

}

