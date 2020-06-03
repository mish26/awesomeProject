/*
test library
 */
package test

import "fmt"

// AnimalAnimalDoc
type AnimalAnimal struct {
	// AnimalName
	Name string
	// AnimalAge
	Age int
}

func (p *AnimalAnimal) Say(){
	fmt.Println("AnimalAnimalWanWan")
}


// average returns Intermediate value of a set of numbers
func Average(s []int)int {
	total := 0
	for _, i := range s{
		total += i
	}
	return int(total/len(s))
}

