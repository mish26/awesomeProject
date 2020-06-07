package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func goroutine(ctx context.Context) {
	// TryAcquireで設定した数のgoroutineのみ処理を走らせることができる
	// 2つめ以降の処理はreturnで走らせない
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("could not get lock")
		return
	}

	// Acquireでlockする
	//if err := s.Acquire(ctx, 1); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	defer s.Release(1)

	fmt.Println("run")
	time.Sleep(1 * time.Second)
	fmt.Println("finish")
}

func main(){
	ctx := context.TODO()
	go goroutine(ctx)
	go goroutine(ctx)
	go goroutine(ctx)
	time.Sleep(2 * time.Second)
	go goroutine(ctx)
	time.Sleep(5 * time.Second)

}