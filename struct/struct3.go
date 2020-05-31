package main

import "fmt"

type Square struct {
	// 小文字にすると同じパッケージ内からしかアクセスできない
	x int
	y int
}

// ポインタレシーバー
// 実体のfieldに対し、値を上書きしている
func (t *Square) Scale(i int) {
	t.x = t.x * i
	t.y = t.y * i
}

// オブジェクト指向っぽい振る舞いになる
// 値レシーバー
func (t Square) Area() int {
	return t.x * t.y
}

// コンストラクタっぽい振る舞い
func New(x, y int) *Square {
	return &Square{x,y}
}

func main() {
	// t := Square{3, 4}
	// fmt.Println(AreaCalculation(t))
	s := New(3, 5)
	s.Scale(10)
	fmt.Println(s.Area())

}

