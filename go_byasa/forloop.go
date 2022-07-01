package main

import (
	"fmt"
)

func main2() {

	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{12, 13, 14, 15, 16}

	c := [][]int{a, b}

	fmt.Println(c)
	fmt.Println(c[1][2])

}
