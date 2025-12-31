package main

/*
使用两个Goroutine，向标准输出中按顺序按顺序交替打出字母与数字，输出是a1b2c3
*/
import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	s := "a1b2c3"
	dch := make(chan uint8, 1)
	sch := make(chan uint8, 1)

	s_send := make(chan struct{}, 1)
	d_send := make(chan struct{}, 1)

	go func() {
		for i := 0; i < len(s); i++ {
			if '0' <= s[i] && s[i] <= '9' {
				dch <- s[i]
			} else if 'a' <= s[i] && s[i] <= 'z' {
				sch <- s[i]
			}
		}
		close(sch)
		close(dch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range sch {
			<-s_send
			fmt.Println(string(n))
			d_send <- struct{}{}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for n := range dch {
			<-d_send
			fmt.Println(string(n))
			s_send <- struct{}{}
		}
	}()

	s_send <- struct{}{}

	wg.Wait()
}
