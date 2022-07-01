package main

import (
	"fmt"
)

func appendString() func(string) string {
	t := "Hello"
	fmt.Println("The t inside appendString:", t)
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func closure1() {
	a := appendString()
	b := appendString()

	fmt.Println(a("MCD"))
	fmt.Println(b("Byasadev"))

	fmt.Println(a("Kusunpur"))
	fmt.Println(b("Nayak"))
}
