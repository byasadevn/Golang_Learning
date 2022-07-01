package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func readfile2() {
	fptr := flag.String("fname", "hello1.go", "Provide the file name")
	flag.Parse()
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	text := bufio.NewScanner(f)
	for text.Scan() {
		fmt.Println(text.Text())
	}

	err = text.Err()
	if err != nil {
		log.Fatal(err)
	}

}
