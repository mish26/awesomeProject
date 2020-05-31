package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)
	f32 := float32(x)
	fmt.Printf("%T %v %f\n", f32, f32, f32)
	var s string = "15"
	fmt.Printf("%T %v %f\n",s, s, s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("%T %v %d\n",i, i, i)
	j, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d\n",j, j, j)

	var a [2]int
	a[0] = 100
	a[1] = 500
	fmt.Println(a)

	var c [2]int = [2]int{100, 200}
	fmt.Println(c)
	var e []int = []int{100, 200, 300}
	e = append(e, 10000, 2000)
	fmt.Println(e, e[1], e[1:3], e[1:2], e[:2], e[2:])
	fmt.Println(e[:])

	var twoDimensional = [][]int{
		[]int{100,200,300},
		[]int{1000,2000,3000},
	}
	fmt.Println(twoDimensional)

	// len => 長さ
	// cap => キャパシティ。メモリ領域に、確保する値
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 10, 20, 30)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	// 出力結果は一緒だが、メモリに領域を確保の有無の差がある
	// lは領域確保する、oはしない(nilなので)。
	l := make([]int, 0)
	var o []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(l), cap(l), l)
	fmt.Printf("len=%d cap=%d value=%v\n", len(o), cap(o), o)

	// 長さ5, キャパシティ5
	t := make([]int, 5)
	for z := 0; z < 5; z++{
		t = append(t, z)
		fmt.Println("t:", t)
	}
	fmt.Println("t:", t)

	fmt.Println()

	// 長さ0, キャパシティ5
	u := make([]int, 0, 5)
	for z := 0; z < 5; z++{
		u = append(u, z)
		fmt.Println("u:", u)
	}
	fmt.Println("u:", u)
}


