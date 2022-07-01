package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func readfile1() {
	test2()
}

func test2() {
	fptr := flag.String("fflag", "hello.go", "file path")
	flag.Parse()
	fmt.Println("The value of the fflag", *fptr)

}

func test1() {

	data, err := ioutil.ReadFile("hello1.go")
	if err != nil {
		fmt.Println("Error in the file reading", err)
		return
	}
	fmt.Println(string(data))
}
