package main

import "fmt"

type Employee struct {
	name string
	id   int
}

type Engineer struct {
	name   string
	id     int
	salary int
}

type Display interface {
	get_name() string
}

func (e *Employee) set_name(name string) {
	e.name = name
}

func (e *Employee) set_id(id int) {
	e.id = id
}

func (e *Engineer) set_name(name string) {
	e.name = name
}

func (e *Engineer) set_id(id int) {
	e.id = id
}

func (e *Engineer) set_salary(salary int) {
	e.salary = salary
}

func (e Engineer) get_name() string {
	return e.name
}

func (e Employee) get_name() string {
	return e.name
}

func main() {
	em := Employee{}
	en := Engineer{}
	em.set_name("Bapan")
	em.set_id(12)
	en.set_name("Ritu")
	en.set_id(33)
	en.set_salary(500)
	fmt.Println(em)
	fmt.Println(en)
	employee := []Display{em, en}
	for _, obj := range employee {
		fmt.Println(obj.get_name())
	}
}
