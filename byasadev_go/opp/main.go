package main

import (
	"opp/employee"
)

func main() {
	e1 := employee.Employee{
		FirstName:   "byasadev",
		LastName:    "nayak",
		TotalLeaves: 30,
		LeaveTaken:  10}
	e1.RemainingLeaves()

}
