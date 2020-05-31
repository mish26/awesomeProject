package main

import "fmt"

type Triangle struct {
	X int
	Y int
}

// ポインタレシーバー
// 実体のfieldに対し、値を上書きしている
func (t *Triangle) Scale(i int) {
	t.X = t.X * i
	t.Y = t.Y * i
}

// オブジェクト指向っぽい振る舞いになる
// 値レシーバー
func (t Triangle) Area() int {
	return t.X * t.Y
}

func Area(t Triangle) int {
	return t.X * t.Y
}

func main() {
	t := Triangle{3, 4}
	fmt.Println(Area(t))
	fmt.Println(t.Area())
	t.Scale(10)
	fmt.Println(t.Area())
}

