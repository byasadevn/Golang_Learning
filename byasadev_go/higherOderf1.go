package main

import (
	"fmt"
)

func simple(a func(b int, c int) int) {
	fmt.Println(a(60, 70))
}

func HigerOrder1() {
	a := func(b, c int) int {
		return b + c
	}
	simple(a)
}
