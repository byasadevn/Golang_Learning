package main

import (
	"fmt"
)

type MyString string

func interface1() {
	ms := MyString("byasadevnayak")
	var v MyVowel
	v = ms
	fmt.Printf("The value: %c", v.countVowel())
}

type MyVowel interface {
	countVowel() []rune
}

func (ms MyString) countVowel() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}
