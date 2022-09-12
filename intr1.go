package main

import "fmt"

type Human interface {
	get_name() string
	get_id() int
}

type Person struct {
	name string
	id   int
}

func (p Person) get_name() string {
	return p.name
}

func (e *Employee) get_name() string {
	return e.name
}

func (p Person) get_id() int {
	return p.id
}

func (e *Employee) get_id() int {
	return e.id
}

type Employee struct {
	name string
	id   int
}

func main() {
	var h1 Human
	p1 := Person{"Bapan", 12}
	h1 = p1
	em := Employee{"Bapan", 13}
	fmt.Println(h1.get_name())
	fmt.Println(h1.get_id())
	h1 = &em
	fmt.Println(h1.get_name())
	fmt.Println(h1.get_id())
}
