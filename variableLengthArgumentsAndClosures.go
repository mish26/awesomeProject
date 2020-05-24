package main

import "fmt"

func incrementGenerator() (func() int) {
	x := 0
	return func() int{
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func variadicArguments(args ...int)  {
	fmt.Println(len(args), cap(args), args)
	for _, arg := range args{
		fmt.Println(arg)
	}
}

/*
	クロージャと可変長引数のsyntax確認
 */
func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	cA1 := circleArea(3.1415)
	fmt.Println(cA1(2))

	cA2 := circleArea(3)
	fmt.Println(cA2(2))

	variadicArguments(10,2,400,500)
	variadicArguments()

	// 可変長引数のfunctionにsliceを適用する
	s := []int{1,2,3}
	fmt.Println(s)
	variadicArguments(s...)
}


