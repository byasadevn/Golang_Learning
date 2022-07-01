package main

import (
	"fmt"
	"os"
)

func main() {

	//var wer interface{}  = 23
	fmt.Println(wer)

	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString("Hello to Go lang file handling\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes are written successfully")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
