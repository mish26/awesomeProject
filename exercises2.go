package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	a := 0
	for i, value := range l {
		if i == 0 {
			a = value
			continue
		}

		if a > value {
			a = value
		}
	}
	fmt.Println(a)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _, value := range m{
		sum += value
	}
	fmt.Println(sum)

}


