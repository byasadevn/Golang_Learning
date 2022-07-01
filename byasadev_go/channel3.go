package main

import "fmt"

func findDigit(data int, ch chan int) {
	for data != 0 {
		ch <- data % 10
		data /= 10
	}
	close(ch)
}

func calSquare(data int, ch1 chan int) {

	localch := make(chan int)
	go findDigit(data, localch)
	sum := 0
	for v := range localch {
		sum += v * v
	}
	ch1 <- sum
}

func calCube(data int, ch2 chan int) {
	localch := make(chan int)
	go findDigit(data, localch)
	sum := 0
	for v := range localch {
		sum += v * v * v
	}
	ch2 <- sum
}

func channel3() {
	val := 589
	mch1 := make(chan int)
	go calSquare(val, mch1)
	//asd := <-mch1
	//fmt.Println(asd)
	mch2 := make(chan int)
	go calCube(val, mch2)

	fmt.Println(<-mch1 + <-mch2)
}
