package main

import (
	"fmt"
)

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Cat struct {
	Name string
}

func (p *Person) Say() string{
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	fmt.Printf("%T %v\n", mike, mike)
	DriveCar(mike)
	var taro Human = &Person{"taro"}
	DriveCar(taro)
	// cannot use cat (type Cat) as type Human in argument to DriveCar:
	//	Cat does not implement Human (missing Say method)
	// 上記エラーが出る。DriveCarは、Humanの型を実装したものしかうけつけていない
	// var cat Cat = Cat{"nyanko"}
	// DriveCar(cat)
}

