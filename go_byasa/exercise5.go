package main

import (
	"fmt"
)

func execrcise5() {

	fmt.Println("Inside the main method")

	fmt.Println("-------------------------------------------------------------------")
	//------------------------------------
	a1 := []int{1, 2, 3, 4}
	b1 := make([]int, 2)
	b1[0] = 6
	b1[1] = 7

	//a1 = b1
	//copy(a1, b1)
	a1 = append(a1, b1...)

	fmt.Println(a1)
	fmt.Println(b1)

	//---------------------------------------
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	fmt.Printf("type of a :%T\n", a)
	//------------------------------------
	s1 := a[3:5]
	fmt.Printf("type of s1 :%T, len:%v and cap:%v\n", s1, len(s1), cap(s1))
}
