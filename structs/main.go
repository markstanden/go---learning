package main

import (
	"fmt"
)

func main() {
	jim := person{
		first:  "Jim",
		second: "Earthworm",
		contact: contactInfo{
			email:       "Jim@jim.com",
			phoneNumber: 12432512341234,
		},
	}
	fmt.Printf("%+v\n", jim)
}
