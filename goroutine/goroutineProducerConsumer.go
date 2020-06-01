package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(c chan int, i int) {
	c <- i * 2
}

func consumer(ch chan int, waitGroup *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i * 1000)
		waitGroup.Done()
	}
	fmt.Println("-----------------")
}

func main() {
	var waitGroup sync.WaitGroup
	chanInt := make(chan int)

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go producer(chanInt, i)
	}

	go consumer(chanInt, &waitGroup)
	waitGroup.Wait()
	close(chanInt)
	time.Sleep(2)
	fmt.Println("close後")



}



