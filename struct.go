package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
	S string
	A bool
	T, D int
}

type Hoge struct {
	hogeInt int
	hogeInt2 int
	hogeStr string
}

func changeHoge(hoge Hoge)  {
	hoge.hogeInt = 1000
}
func changeHoge2(hoge *Hoge)  {
	//　(*hoge).hogeInt = 1000
	// 上記の省略形
	// structの場合は、引数が実態の場合は代入するfieldも実態を指す
	hoge.hogeInt = 1000


}

func main() {
	fuga := Hoge{1, 2, "sample"}
	changeHoge(fuga)
	fmt.Println(fuga)

	fuga2 := &Hoge{1, 2, "sample"}
	changeHoge2(fuga2)
	fmt.Println(fuga2)
	fmt.Println(*fuga2) // pointerの実態を出力

	v := Vertex{X:1, Y:2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X:1}
	fmt.Println(v2)

	// keyを指定しない場合は、fieldの数分、引数を設定する必要がある。
	// 型も合わせる必要がある
	v3 := Vertex{1,2,"test", true, 1, 0}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{} // pointerを返すときは、&を使う人が多い。アドレスが明示的でわかりやすい。
	fmt.Printf("%T %v\n", v7, v7)

	// sliceは、makeを使うほうが良い
	slice := make([]int, 0)
	slice2 := []int{}
	fmt.Println(slice)
	fmt.Println(slice2)


}

