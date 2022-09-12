package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_string(str string) string {
	fmt.Println("Enter ", str)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gender := scanner.Text()
	return gender
}

func main() {
	gender := get_string("gender")
	name := get_string("name")
	if gender == "male" {
		fmt.Println("Your name is :", name)
		fmt.Println("Your gender is:", gender)
	} else if gender == "female" {
		fmt.Println("Your name is:", name)
		fmt.Println("Your gender is:", gender)
	} else {
		fmt.Println("Are you a transgender:?", name)
	}
}
