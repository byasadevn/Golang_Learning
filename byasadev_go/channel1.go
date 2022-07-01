package main

import "fmt"

func QuareSum(num int, sqch chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	//fmt.Printf("The vale: %v", sum)
	sqch <- sum

}

func QuadSum(num int, quch chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	//fmt.Printf("The vale: %v", sum)
	quch <- sum
}

func test1(mych chan<- int) {
	mych <- 10
}

func channel1() {

	wer := make(chan int)
	go test1(wer)
	fmt.Printf("The value= %v\n", <-wer)

	fmt.Println("-------------------------------------")

	digit := 123
	sqch := make(chan int)
	quch := make(chan int)

	go QuadSum(digit, sqch)
	go QuareSum(digit, quch)
	//<-sqch
	//<-quch
	sqsum, qusum := <-sqch, <-quch
	total := sqsum + qusum
	fmt.Printf("The value %v", total)

}
