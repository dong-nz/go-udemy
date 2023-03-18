package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	address string
	phone   string
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
