package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string, 40)

	wg.Add(1)
	go func(ch1 chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 40; i++ {
			ch1 <- strconv.Itoa(i)
		}
		close(ch1)
	}(ch1, &wg)

	wg.Add(1)
	go func(ch1 <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			num, ok := <-ch1
			if ok == true {
				fmt.Println("Routine 1: " + num)
				time.Sleep(1 * time.Millisecond)
			}
		}
	}(ch1, &wg)

	wg.Add(1)
	go func(ch1 <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			num, ok := <-ch1
			if ok == true {
				fmt.Println("Routine 2: " + num)
				time.Sleep(2 * time.Millisecond)
			}
		}
	}(ch1, &wg)

	wg.Wait()

}
