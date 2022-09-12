package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	subjects := make(chan string, 5)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			subjects <- "subject " + strconv.Itoa(i)
		}
		close(subjects)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			su, ok := <-subjects
			if ok == true {
				fmt.Println(su)
			}
		}
	}(&wg)

	wg.Wait()

}
