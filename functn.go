package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	name string
	id   int
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) setid(id int) {
	e.id = id
}

func (e Employee) display() string {
	id := strconv.Itoa(e.id)
	return id + " " + e.name
}

func main() {
	em := Employee{}
	em.setName("Bapan")
	em.setid(12)
	fmt.Println(em.display())
}
