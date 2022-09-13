package main

import (
	"fmt"
	"os"
)

func main() {
	contents, err := os.ReadFile("./test1.txt")
	if err != nil {
		fmt.Println("Error occured:")
		return
	}
	fmt.Println("The contents of the file is:", string(contents))
}
