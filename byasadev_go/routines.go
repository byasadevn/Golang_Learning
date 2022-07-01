package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello Python")
}

func printChar() {
	for i := 'a'; i < 'e'; i++ {
		fmt.Printf("The char: %c\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printNumber() {
	for i := 0; i < 10; i++ {
		fmt.Printf("The number: %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func routines123() {

	fmt.Println("-------Entry main-----------")
	go hello()
	go printChar()
	go printNumber()
	time.Sleep(2 * time.Second)
	fmt.Println("-------Exit main------------")

	var a chan int
	if a == nil {
		fmt.Println("The chanell a is nil")
		a = make(chan int)
		fmt.Println(a)
	}

}
