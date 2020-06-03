package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	breakLoop:
		for {
			select {
			case t := <-tick:
				fmt.Println("tick.", t)
			case <-boom:
				fmt.Println("BOOM!")
				// return
				// breakでどこのループを抜けるのかを定義できる
				break breakLoop
			default:
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
			}
		}
	fmt.Println("----------")
}
