package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func get_number() int64 {
	fmt.Println("Enter the number:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return num
}

func main() {
	num1 := get_number()
	num2 := get_number()
	num3 := num1 + num2
	fmt.Println("The sum of the number is:", num3)
}
