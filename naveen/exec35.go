package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to exec35")
	//test1()
	test2()
}

func test2() {
	const size = 5
	f, err := os.Open("text1.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	buff := make([]byte, size)

	for {
		_, err := f.Read(buff)

		if err != nil || err == io.EOF {
			break
		}

		fmt.Println(string(buff))

	}

}
func test1() {
	f, err := os.Open("text1.txt")
	if nil != err {
		fmt.Println("Unbale to open the file in read mode", err)
		return
	}
	defer f.Close()
	fmt.Println("Suceessfully opened file", f.Name())
	//data := fmt.Fscan(f)
	//fmt.Println(string(data))
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
