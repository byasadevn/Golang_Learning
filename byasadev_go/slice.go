package main

import (
	"fmt"
)

func slice() {
	slice2()
	//slice3()
}
func slice3() {
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
}

func slice2() {
	a := []string{"Hello", "World"}
	fmt.Println(a)
	slice1(a...)
	fmt.Println(a)
	//append(a, string("Byasa"))

}

func slice1(s ...string) {
	s[0] = "Go"
	s = append(s, "Playground")

}
