package main

import (
	"fmt"
	"sync"
	"time"
)

func normalOutput(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func parallelOutput(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}
func parallelOutput2(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
    // 並列させたい処理一つにつき一つaddする
	wg.Add(2)
	go parallelOutput("parallelOutput", &wg)

	go parallelOutput2("parallelOutput2", &wg)
	normalOutput("nomal")
	// wgがdoneするまで待つ
	wg.Wait()
}


