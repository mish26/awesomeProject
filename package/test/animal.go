package test

import "fmt"

var PublicVal string = "Public"
var privateVal string = "private"

type Animal struct {
	// AnimalName
	Name string
	// AnimalAge
	Age int
}

type animal2 struct {
	Name string
	Age int
}

func Say() {
	fmt.Println("wanwan")
}

