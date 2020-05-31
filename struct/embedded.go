package main

import "fmt"

type Pentagon struct {
	X int
	Y int
}

// ポインタレシーバー
// 実体のfieldに対し、値を上書きしている
func (t *Pentagon) Scale(i int) {
	t.X = t.X * i
	t.Y = t.Y * i
}

// オブジェクト指向っぽい振る舞いになる
// 値レシーバー
func (t Pentagon) Area() int {
	return t.X * t.Y
}

// embedded
// 継承っぽい振る舞いができる
type Pentagon3D struct {
	Pentagon
	z int
}

// ポインタレシーバー
// 実体のfieldに対し、値を上書きしている
func (t *Pentagon3D) Scale3D(i int) {
	t.X = t.X * i
	t.Y = t.Y * i
	t.z = t.z * i
}

// オブジェクト指向っぽい振る舞いになる
// 値レシーバー
func (t Pentagon3D) Area3D() int {
	return t.X * t.Y *t.z
}


// コンストラクタっぽい振る舞い
func new(x, y, z int) *Pentagon3D {
	return &Pentagon3D{Pentagon{x, y}, z}
}

func main() {
	t := new(3, 4, 5)
	fmt.Println("Patagonia3D: ", t)
	t.Scale(10)
	fmt.Println("Patagonia3D: ", t)
	fmt.Println(t.Area())
	fmt.Println(t.Area3D())
}

