package main

import (
	"fmt"
)

type auther struct {
	firstname string
	lastname  string
	bio       string
}

func (au auther) fullname() string {

	return fmt.Sprintf("%s %s", au.firstname, au.lastname)
}

type blog struct {
	title string
	auther
}

func (bl blog) printIndeatils() {
	fmt.Printf("%s %s %s", bl.fullname(), bl.bio, bl.title)
}

func composition() {
	a := auther{"byasadev", "nayak", "golang"}
	b := blog{"Learning abt golang composition feature", a}
	b.printIndeatils()
}
