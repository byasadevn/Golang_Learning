package main

import (
	"fmt"
)

func learnArray() {

	//array1()
	//array2()
	multiArray()
}

func multiArray() {
	a := [3][2]string{
		{"basu", "anita"},
		{"basi", "kandei"},
		{"dunga", "kuni"},
	}
	fmt.Println(a)
}

func array2() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v ", a[i])
	}
	fmt.Printf("\n")
	for i, v := range a {
		fmt.Printf("%v---%v ", i, v)
	}
	fmt.Println()
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
}

func array1() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	b := [3]int{4, 5, 6}
	fmt.Println(b)

	c := [4]float64{12.2}
	fmt.Println(c)

	d := [...]float32{1.2, 2.3, 3.4}
	fmt.Println(d)

}
