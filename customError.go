package main

import "fmt"

func myFunc() error {
	isNormal := false
	if isNormal {
		return nil
	}
	return &AnimalNotFound{name: "wanko"}
}

// 独自エラー
type AnimalNotFound struct {
	name string
}

func (e *AnimalNotFound) Error() string {
	return fmt.Sprintf("animal not found: %v", e.name)
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}

