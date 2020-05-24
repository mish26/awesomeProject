package main

import (
	"fmt"
	"runtime"
)

func main() {
	okFunk := func() string {
		return "OK"
	}

	result := okFunk()
	if result == "OK" {
		fmt.Println("result true")
	}

	// 上記のif文を省略して書くやり方
	if result2 := okFunk(); result2 == "OK" {
		fmt.Println("result2 true")
	}

	s := []string{"te", "tes", "test"}
	for i := 0; i < len(s); i++{
		fmt.Println(i, s[i])
	}
	// 上記のfor文をrangeで実装
	for i, v := range s {
		fmt.Println(i, v)
	}
	for _, v := range s {
		fmt.Println(v)
	}

	// map
	m := map[string]int{"go":80, "java":60}
	for key, value := range m{
		fmt.Println(key, value)
	}

	for key := range m{
		fmt.Println(key)
	}
	for _,value := range m{
		fmt.Println(value)
	}

	// switch
	os := getOsName()
	fmt.Println(os)
	switch os {
	case "darwin":
		fmt.Println("Mac!OS")
	case "windows":
		fmt.Println("windows!")
	default:
		fmt.Println("Default!")
	}

	// 省略形
	switch os2 := getOsName(); os2 {
	case "darwin":
		fmt.Println("Mac!OS2", os2)
	case "windows":
		fmt.Println("windows!")
	default:
		fmt.Println("Default!")
	}

	// 省略形の変数はスコープ内でしか使えない
	// fmt.Println(os2)

	// switchの条件書かなくても実装できる
	os3 := getOsName()
	switch {
	case os3 == "darwin":
		fmt.Println("Mac!OS3", os3)
	case os3 == "windows":
		fmt.Println("windows!")
	default:
		fmt.Println("Default!")
	}
}

func getOsName() string {
	return 	runtime.GOOS
}


