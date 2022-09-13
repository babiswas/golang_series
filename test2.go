package main

import "fmt"

func add(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := add(nums...)
	fmt.Println("The sum of the numbers is:", sum)
	sum = add(5, 6, 7, 8, 9)
	fmt.Println("The sum of the numbers is:", sum)
}
