package main

import (
	"fmt"
)

func firstclassfunc() {

	a := func() {
		fmt.Println("Hello from the first class function")
	}
	a()
	fmt.Printf("The type of a:%T", a)
}
