package main

import "fmt"

type Person struct {
	name       string
	id         int
	occupation string
}

func main() {
	var p Person
	var name string
	var id int
	var occupation string
	fmt.Println("Enter the name:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Enter the id:")
	fmt.Scanf("%d\n", &id)
	fmt.Println("Enter the occupation:")
	fmt.Scanf("%s\n", &occupation)
	p.name = name
	p.id = id
	p.occupation = occupation
	fmt.Println(p)
}
