package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fptr := flag.String("filename", "test1.txt", "file name for reading")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			fmt.Println("Finished reading the file:")
			break
		}
		if err != nil {
			fmt.Printf("Error %s reading the file", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}
}
