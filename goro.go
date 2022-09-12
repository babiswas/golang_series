package main

import (
	"fmt"
	"sync"
	"time"
)

func test(str1 string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(str1)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go test("hello", &wg)
	wg.Add(1)
	go test("bello", &wg)
	wg.Wait()
}
