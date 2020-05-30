package main

import "fmt"

/*
どんな型でも受け付ける
JavaのObject型と同じ
 */
func do(i interface{})  {
	// 下記はエラー。iはinterface型で、キャストする必要がある
	// i2 := i * 2

	// type assertion
	// interfaceを型に変換することはtype assertionという
	// 型を、紐づく別の型に変換することは、type conversion
	i2 := i.(int)
	i2 *= 2
	fmt.Println(i2)
}

func do2(i interface{})  {
	ss := i.(string)
	fmt.Println(ss + "!")
}

func doSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("default %T %v\n", v, v)
	}
}


func main() {
	do(10)
	do2("Mikeda")
	doSwitch(20)
	doSwitch("Test")
	doSwitch(false)

	// interfaceではない、intからfloatに型変換は、type conversion
	// type conversion(キャスト)
	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}

