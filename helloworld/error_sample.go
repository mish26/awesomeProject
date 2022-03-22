package main

import (
	"log"
)

type MyError1 struct {}

func (e *MyError1) Error() string {
	return "my error 1"
}

type MyError2 struct {}

func (e *MyError2) Error() string {
	return "my error 2"
}

func main() {
	if err := doError(1); err != nil {
		switch e := err.(type) {
		case *MyError1:
			log.Fatalf("%#v", e)
		case *MyError2:
			log.Fatalf("%#v", e)
		}
	}
}

func doError(i int) error {
	switch i {
	case 1:
		return &MyError1{}
	case 2:
		return &MyError2{}
	}
	return nil
}
