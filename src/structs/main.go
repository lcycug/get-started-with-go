package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contact: contact{
			email:   "jsmith@gmail.com",
			zipCode: 9000,
		},
	}
	jim.updateFirstName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
