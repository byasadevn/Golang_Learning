package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func exec36() {
	fmt.Println("Welcome to exe36")
	//test1()
	test2()
}

func test2() {

	val := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < 300; i++ {
		wg.Add(1)
		go generateRand(val, &wg)
	}

	go consumeData(val, done)

	go func() {
		wg.Wait()
		close(val)

	}()

	ret := <-done
	if ret == true {
		fmt.Println("Successfully written")
	} else {
		fmt.Println("Not able to write to file")
	}

}

func generateRand(val chan int, wg *sync.WaitGroup) {
	i := rand.Intn(999)
	val <- i
	fmt.Println("Inside the rand")
	wg.Done()
}
func consumeData(val chan int, done chan bool) {

	f, err := os.Create("text2.txt")

	if nil != err {
		done <- false
		fmt.Println("Got the error while opening the file", err)
		return
	}
	for i := range val {
		fmt.Println("Inside range")
		fmt.Fprintln(f, i)

	}

	err = f.Close()
	if nil != err {
		done <- false
		fmt.Println("File closing error", err)
		return
	}

	done <- true
}

func test1() {
	//f, err := os.Create("text1.txt")
	f, err := os.OpenFile("text1.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if nil != err {
		fmt.Println("File opening error", err)
		return
	}
	/*l, err := f.WriteString("Welcome to go file writing")
	if nil != err {
		f.Close()
		fmt.Println("File writing error: ", err)
		return
	}*/
	/*text := []byte{100, 101, 102, 103, 104, 105, 106, 107, 108, 109}
	l, err := f.Write(text)
	*/
	multiline := [...]string{"Today is Sunday", "We have a lot plan", "But, we able to do few at the end"}
	for _, line := range multiline {
		fmt.Fprintln(f, line)
	}
	fmt.Fprintln(f, "The new line appended")
	/*if nil != err {
		f.Close()
		fmt.Println("File writing error: ", err)
		return
	}*/
	//fmt.Println("The number of bytes written :", l)
	fmt.Printf("Type of multiline %T", multiline)
	err = f.Close()
	if nil != err {
		fmt.Println("Unable to close the file :", err)
		return
	}

}
