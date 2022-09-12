package main

import "fmt"

func get_marks() [5]int {
	var marks [5]int
	for i := 0; i < 5; i++ {
		fmt.Println("Enter the mark:")
		fmt.Scanf("%d\n", &marks[i])
	}
	return marks
}

func add_marks(arr [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	sum := add_marks(get_marks())
	fmt.Println("The total marks obtained by the candidate is:", sum)
	return
}
