package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 10)

	go gen(ch, 0)
	go gen(ch, 100)
	go func() {
		for {
			v := <-ch
			fmt.Println(v)
		}
	}()
	time.Sleep(10 * time.Second)
}

func gen(ch chan string, start int) {
	for i := 0 + start; i < 10+start; i++ {
		select {
		case ch <- fmt.Sprintf("%d", i):
			// time.Sleep(100)
		default:
			fmt.Println("...")
		}
	}
}
