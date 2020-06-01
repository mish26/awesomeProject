package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string){
	for {
		ch <- "connect 1"
		time.Sleep(1+time.Second)
	}
}
func goroutine2(ch chan string) {
	for {
		ch <- "connect 2"
		time.Sleep(4 + time.Second)
	}
}

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	// 複数のAPIを叩くときに使うイメージ
	go goroutine1(c1)
	go goroutine2(c2)

	for {
		select {
		// 同時に複数の関数を並列に叩くことができる
		// c2でレスポンス4sだが、c2のレスポンスを待つことなく、処理を進めることができる
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}



}