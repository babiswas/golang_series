package main

import "fmt"

func main() {
	m := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println("The sum of the number is:", m)
}
