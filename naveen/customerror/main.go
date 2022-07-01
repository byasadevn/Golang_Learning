package main

import (
	"errors"
	"fmt"
	"math"
)

func main123() {

	area, err := calArea(-22.3)
	if err != nil {
		fmt.Println("there is an error : -", err)
		return
	}
	fmt.Printf("The are : %0.2f", area)

}

func calArea(rad float64) (float64, error) {

	if rad < 0 {
		return 0, errors.New("Radius less than zero")
	}

	return math.Pi * rad * rad, nil
}
