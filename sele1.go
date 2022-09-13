package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)

	wg.Add(1)
	go func(ch1 chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			if i%2 != 0 {
				ch1 <- i
			}
		}
		close(ch1)
	}(ch1, &wg)

	go func(ch1 chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				ch2 <- i
			}
		}
		close(ch2)
	}(ch2, &wg)

	for i := 0; i < 100; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
