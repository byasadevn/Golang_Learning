package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func exec30() {
	//test1()
	//test2()
	test3()
}

var ErrBadPattern = errors.New("Basudev")

func test3() {
	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("bad pattern :", err)
			return
		}
		fmt.Println("Generic err: ", err)
		return
	}
	fmt.Println("The files: ", files)

}

func test2() {

	add, err := net.LookupHost("golangboot.com")
	if err != nil {
		if aErr, ok := err.(*net.DNSError); ok {

			if aErr.Timeout() {
				fmt.Println("Time out error", aErr)
				return
			}
			if aErr.Temporary() {
				fmt.Println("Temporary error", aErr)
				return
			}
			fmt.Println("DNS error", aErr)
			return
		}
		fmt.Println("Generic error,err")
		return
	}
	fmt.Println(add)

}

func test1() {
	f, err := os.Open("text3.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Printf("The path err: %s", pErr.Path)
			return
		}
		fmt.Println("Generic error: ", err)
	}
	fmt.Println("The file opened successfully: ", f.Name())
}
