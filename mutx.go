package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	ch1 := make(chan string)
	sea := []string{"gold", "silver", "diamond", "ruby", "pearl", "iron", "copper", "jewellery", "brass"}
	storehouse := make([]string, 4)

	store_wealth := func(wg *sync.WaitGroup, ch <-chan string) {
		defer wg.Done()
		wealth, ok := <-ch
		if ok == True {
			fmt.Println("Storing wealth in the store:")
			storehouse = append(storehouse, wealth)
		}
	}

	go func(wg *sync.WaitGroup, mu *sync.Mutex, ch chan<- string) {
		defer wg.Done()
		mu.Lock()
		wlth := sea[rand.Intn(9)]
		mu.Unlock()
		ch1 <- wlth
	}(&wg, &mut, ch1)

}
