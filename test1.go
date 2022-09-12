package main

import "fmt"

func find_type(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("This is a string type:")
	case int:
		fmt.Println("This is an integer type:")
	default:
		fmt.Println("This is default block:")
	}
}

func main() {
	find_type("hello")
	find_type(21)
	find_type("bello")
	find_type(true)
}
