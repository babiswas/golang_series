package main

import (
	"flag"
	"fmt"
)

func main() {
	fptr := flag.String("fpath", "test1.txt", "file path to read from")
	flag.Parse()
	fmt.Println("Value of fath is:", *fptr)
	numbers := [5]int{1, 2, 3, 4, 5}
	mynumbers := numbers[2:5]
	fmt.Println(mynumbers)
}
