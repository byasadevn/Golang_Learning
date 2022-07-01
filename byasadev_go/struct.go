package main

import (
	"fmt"
)

func strcu01() {

	//struct1()
	//struct2()
	//struct33()
	//struct412()
	struct5()

}

func struct5() {
	type Person struct {
		string
		int
	}
	p1 := Person{
		string: "naveen",
		int:    25,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}

func struct412() {

	type qaz struct {
		name string
		age  int
	}

	rfv := &qaz{"ert", 24}
	fmt.Println(*(rfv))

}

func struct2() {
	emp1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "basu",
		lastName:  "mcd",
		age:       42,
	}
	fmt.Println(emp1)
}

func struct1() {
	type employee struct {
		firstName string
		lastName  string
		age       uint
	}
	emp1 := employee{"byasadev", "nayak", 42}
	fmt.Println(emp1)

	var emp2 employee
	emp2.firstName = "Aradhya"
	emp2.lastName = "Nayak"
	emp2.age = 9
	fmt.Println(emp2)

}

func struct33() {

	emp4 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "basi",
		lastName:  "raju",
		age:       12,
	}

	fmt.Println("The emp4: ", emp4)

}
