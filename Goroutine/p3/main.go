package main

import (
	"fmt"
	"sync"
	"time"
)

/*
编写一个程序限制10个goroutine执行，每执行完一个goroutine就放一个新的goroutine进来
*/
var wg sync.WaitGroup
var count = 1

func main() {

	c := make(chan struct{}, 10)
	for {
		wg.Add(1)
		n := count
		c <- struct{}{}
		go func() {
			defer wg.Done()
			fmt.Printf("执行任务%v\n", (n))
			time.Sleep(1 * time.Second)
			<-c
		}()
		count++
	}
	wg.Wait()

}
