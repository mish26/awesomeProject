package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("%b %T %v", u8, u8, u8)
	fmt.Println()
	fmt.Println("1 + 1 = ", 1+1)
	x := 1 + 1
	fmt.Println(x)
	fmt.Println("2 + 2 =", 2+2)
	fmt.Println("2 + 2 ="[0])

	// アスキーコードで表示される
	fmt.Println("HelloWorld"[0])
	fmt.Println("HelloWorld"[1])

	// 文字列で表示
	fmt.Println(string("HelloWorld"[1]))

	// 文字列置換
	var text string = "Hello World"
	fmt.Println(strings.Replace(text, "o", "Z", 2))
	fmt.Println(strings.Replace(text, "o", "Z", 1))

	// 文字列一致しているかどうか
	fmt.Println(strings.Contains(text, "llo "))

	// 改行出力とダブルクォート出力
	fmt.Println(`
TEST 
 " HELLO worl
d""
`)

	t, f := true, false
	fmt.Printf("%T %v %t\n", t,t,1)
	fmt.Printf("%T %v %t\n", f,f,0)
	fmt.Printf("%T %v %t\n", f,f,f)
	fmt.Println(!t, !f)

}
