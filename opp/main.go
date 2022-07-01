package main

import (
	"opp/employee"
)

func main() {
	e1 := employee.Employee{
		firstName:   "byasadev",
		lastName:    "nayak",
		totalLeaves: 30,
		leaveTaken:  10}
	e1.RemainingLeaves()

}
