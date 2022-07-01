package main

import (
	"fmt"
)

type address struct {
	city  string
	state string
}

type person1 struct {
	firstName string
	lastName  string
	add       address
}

func (a person1) displayAddress() {
	fmt.Printf("The city: %s and state: %s ", a.add.city, a.add.state)
}

func test2() {
	p := person1{
		firstName: "byasadev",
		lastName:  "nayak",
		add: address{
			city:  "kusunpur",
			state: "odisha",
		},
	}
	p.displayAddress()

}

type Employee struct {
	firstName string
	lastname  string
	age       int
}

func method1() {
	//test1()
	test2()
}
