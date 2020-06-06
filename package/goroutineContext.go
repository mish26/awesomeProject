package main

import (
	"context"
	"fmt"
	"time"
)

func goroutine(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main(){
	c1 := make(chan string)
	// goroutineにtimeoutを設定できる
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()
//	ctx := context.TODO()
	go goroutine(ctx, c1)
	// cancel()

	CTXLOOP:
		for {
			select {
			case <- ctx.Done():
				fmt.Println(ctx.Err())
				break CTXLOOP
			case <- c1:
				fmt.Println("success")
				break CTXLOOP
			}
		}
	fmt.Println("~~~~~~~~~~~~~")

}