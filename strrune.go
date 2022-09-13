package main

import "fmt"

func main() {
	s := "hello World"
	fmt.Println(s)
	test := []byte(s)
	fmt.Println(test)
	fmt.Println(string(test))
	fmt.Println(string(s[2]))
	var test1 rune = 'h'
	fmt.Println(string(test1))
	fmt.Printf("%v is the value of s and type is %T \n", s, s)
	fmt.Printf("%v is the value of s and type is %T \n", test1, test1)
}
