package test

import (
	"fmt"
	"testing"
)

func Example() {
	s := []int{10, 20, 30, 40}
	v := Average(s)
	// example
	fmt.Println(v)
}

func ExampleAverage() {
	s := []int{10, 20, 30, 40}
	v := Average(s)
	// aa
	fmt.Println(v)
}

func ExampleAnimalAnimal_Say() {
	animal := AnimalAnimal{"WanWan", 20}
	animal.Say()
}

func TestAverage(t *testing.T) {
	s := []int{10, 20, 30, 40}
	v := Average(s)
	if v != 25 {
		t.Error("Expected 25, got", v)
	}
}
