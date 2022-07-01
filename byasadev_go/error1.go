package main

import (
	"fmt"
	"os"
)

func error1() {

	f, err := os.Open("text.go")
	if err != nil {
		fmt.Println("There is an error to pen the file", err)
		return
	}
	fmt.Printf("The file name :%s", f.Name())
}
