package main

import (
	"fmt"
)

type person struct {
	first   string
	second  string
	married bool
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber int
}

func main() {
	jim := person{
		first:   "Jim",
		second:  "Earthworm",
		married: false,
		contactInfo: contactInfo{
			email:       "Jim@jim.com",
			phoneNumber: 12432512341234,
		},
	}
	jim.updateName("James")
	jim.display()
}

func (p person) display() {
	fmt.Printf("Name: \t\t %v %v\n", p.first, p.second)
	fmt.Printf("Married:\t %v\n", p.married)
	fmt.Printf("Contact email:\t %v\n", p.contactInfo.email)
	fmt.Printf("Contact phone:\t %v\n", p.contactInfo.phoneNumber)
}

func (p *person) updateName (n string) {
	(*p).first = n
}
