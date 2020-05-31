package main

import "fmt"

/*
随時結果を出力する
 */
func goroutineSample3(s []int, c chan int)  {
	sum := 0
	for _, v := range s{
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// queのイメージ
	c := make(chan int, len(s))

	go goroutineSample3(s, c)
	for i := range c {
		fmt.Println(i)
	}

}



