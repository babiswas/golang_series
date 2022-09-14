package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("test2.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newline := "hello world"
	_, err = fmt.Fprintln(f, newline)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File appended sucessfully:")
}
