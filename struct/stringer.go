package main

import "fmt"

/*
// Stringer is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
type Stringer interface {
	String() string
}
 */

type Animal struct {
	Name string
	Age int
}

func (a Animal) String() string{
	return fmt.Sprintf("My Name is %v", a.Name)
}

func main() {
	nyanko := Animal{"Nyanko", 12}
	// stringerによって、出力する内容が上書きされている
	fmt.Println(nyanko)

}

