package main

import "fmt"

// structのような振る舞い
type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 100, 100)
	// iの型は、MyIntで、戻り値の型がintなので、キャストしないとエラーになる
	return int(i * 2)
}
func (i MyInt) Double2() MyInt {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 100, 100)
	// 戻り値の型がMyIntなので、キャストする必要がない
	return i * 2
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
	fmt.Println(myInt.Double2())
}

