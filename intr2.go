package main

import "fmt"

type Human interface {
	work() string
}

type Engineer struct {
	name string
	id   int
	role string
}

func (e *Engineer) set_name(name string) {
	e.name = name
}

func (e *Engineer) set_id(id int) {
	e.id = id
}

func (e *Engineer) set_role(role string) {
	e.role = role
}

func (e *Engineer) work() string {
	return e.role
}

type Teacher struct {
	name string
	id   int
	role string
}

func (t *Teacher) set_name(name string) {
	t.name = name
}

func (t *Teacher) set_id(id int) {
	t.id = id
}

func (t *Teacher) set_role(role string) {
	t.role = role
}

func (t *Teacher) work() string {
	return t.role
}

func get_human(i interface{}) {
	switch v := i.(type) {
	case Teacher:
		fmt.Println(v.work())
	case Engineer:
		fmt.Println(v.work())
	default:
		fmt.Println("Not sure about the type:")
	}
}

func main() {
	t1 := Teacher{"bapan", 12, "Head Teacher"}
	e1 := Engineer{"bapan", 13, "Senior Engineer"}
	fmt.Println(e1)
	fmt.Println(t1)
	get_human(t1)
	get_human(e1)
}
