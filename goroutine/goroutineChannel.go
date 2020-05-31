package main

import "fmt"

func goroutineSample1(s []int, c chan int)  {
	sum := 0
	for _, v := range s{
		sum += v
	}
	// channelに値を渡す
	c <- sum
}

func goroutineSample2(s []int, c chan int)  {
	sum := 1
	for _, v := range s{
		sum *= v
	}
	// channelに値を渡す
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	// queのイメージ
	c := make(chan int)

	go goroutineSample1(s, c)
	go goroutineSample2(s, c)
	// channelから値を受け取る
	// 受信を待つ。waitを実装する必要がない
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

}



