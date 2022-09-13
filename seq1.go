package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	test1 := counter()
	for {
		num := test1()
		fmt.Println("The value of num is :", num)
		if num == 10 {
			break
		}
		fmt.Println(num)
	}
}
