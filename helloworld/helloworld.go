package main

import (
	"fmt"
	"os/user"
	"time"
)


const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

// var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1
// initが最初
func init() {
	fmt.Println("Init")
}

func buzz() {
	fmt.Println("Bazz")
}

func foo()  {
	// Short Variable Declaration
	// 関数の外では実行できない
	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "foo_test"
	xt, xf := true, false
	fmt.Println(xi, xf64,xs,xt,xf)
	fmt.Printf("%T", xi)
}

func main()  {
	// parenthesisで囲める
	var (
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		b, bo bool = true, false
	)
	foo()
	buzz()
	fmt.Println(i,f64,s,b,bo)
	fmt.Println(Pi, Username, Password)
	fmt.Println(user.Current())
	fmt.Println(time.Now())
    fmt.Println(big - 1)
}